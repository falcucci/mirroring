package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// GetFileData(url) retrieves the data from a file given a url.
//
// Parameters:
//
//	url: The url of the file to retrieve.
//
// Returns:
//
//	The data from the file.
func GetFileData(url string) string {
	filename := fmt.Sprintf("./output/%s.html", strings.Replace(url, "/", "_", -1))
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return ""
	}

	fmt.Printf("File %s exists\n", filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return ""
	}
	content := string(data)
	return content
}

// download(url) downloads the content from the specified URL using an
// HTTP GET request and returns it as a string.
//
// Parameters:
// - url (string): The URL from which to download the content.
//
// Returns:
// - string: The content downloaded from the specified URL.
func download(url string) string {
	data := GetFileData(url)
	if data != "" {
		return data
	}

	fmt.Println("Downloading...")
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		println("Error", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("URL: ", url)
		println("Error", err)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println("Error", err)
	}
	return string(respBody)
}

/**
 * ValidateUrl validates a given URL string
 *
 * @param url The URL string to validate
 * @return bool Returns true if the URL is valid, false otherwise
 */
func ValidateUrl(url string) bool {
	fmt.Println("URL: ", url)
	pattern := `^(https?|http)://[^\s/$.?#].[^\s]*$`
	matched, _ := regexp.MatchString(pattern, url)
	return matched
}

// GetHTML downloads the HTML content from the specified URL
// and returns it as a string.
//
// Parameters:
// - url (string): The URL from which to download the HTML content.
//
// Returns:
// - string: The HTML content downloaded from the specified URL.
func GetHTML(url string) string {
	fmt.Println("Getting HTML...")
	html := download(url)
	return html
}

// hasRoutes checks if the specified URL has at least one route.
//
// Parameters:
// - url (string): The URL to check for routes.
//
// Returns:
// - bool: True if the URL has at least one route, false otherwise.
func HasRoutes(url string) bool {
	urlSplitted := strings.Split(url, "/")
	return len(urlSplitted) > 3
}

// GetLinks extracts all links from the given HTML content and
// returns them as a slice of strings.
// The links are extracted using a regular expression pattern
// that matches the href attribute of anchor tags.
//
// Parameters:
// - origin (string): The origin URL of the webpage from which
// the HTML content was extracted.
// - html (string): The HTML content from which to extract links.
//
// Returns:
// - []string: A slice of strings containing all links
// extracted from the HTML content.
func GetLinks(origin string, html string) []string {
	prefix := "/"
	pattern := `href="([^"]+)"`
	originHasRoutes := HasRoutes(origin)

	if originHasRoutes {
		prefix = strings.Split(origin, "/")[3]
		prefix = fmt.Sprintf(`/%s`, prefix)
	}

	// TODO: improve this function to avoid duplicates
	fmt.Println("Getting links...")
	// TODO: modify this pattern to match all links and paths
	// using the same domain
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(html, -1)
	fmt.Println("Matches: ", matches)
	var links []string
	for _, match := range matches {
		url := origin
		path := match[1]
		hasPrefix := strings.HasPrefix(path, prefix)
		hasOrigin := strings.HasPrefix(path, origin)
		if hasPrefix {
			url = origin + path
			fmt.Println("Matched: ", path)
		}

		if hasOrigin {
			url = path
		}

		links = append(links, url)
	}
	return links
}

// SaveFile saves the given HTML content to a file with a filename generated from the URL parameter.
// The file is saved in the specified output directory. Returns true if the file was successfully saved,
// false otherwise.
//
// Parameters:
// - url (string): The URL of the webpage from which the HTML content was extracted.
// - html (string): The HTML content to be saved to a file.
// - output (string): The directory where the file will be saved.
//
// Returns:
// - bool: True if the file was successfully saved, false otherwise.
func SaveFile(url string, html string, output string) bool {
	filename := fmt.Sprintf("%s.html", strings.Replace(url, "/", "_", -1))
	fmt.Println("Saving file...", filename)
	err := os.MkdirAll(output, 0755)
	if err != nil {
		println("Error", err.Error())
		return false
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s", output, filename), []byte(html), os.ModePerm)
	if err != nil {
		println("Error", err.Error())
		return false
	}
	return true
}
