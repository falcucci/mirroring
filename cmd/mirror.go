package cmd

import (
	"fmt"
	"mirroring/utils"
	"sync"

	"github.com/spf13/cobra"
)

var waitGroup sync.WaitGroup

var mirror = &cobra.Command{
	Use:   "mirror",
	Short: "Mirror a URL",
	Long:  `Mirror Web Crawler`,
	Run: func(cmd *cobra.Command, args []string) {
		var url = args[0]
		var output = args[1]
		fmt.Println("URL: ", url)
		fmt.Println("Output: ", output)
		matched := utils.ValidateUrl(url)
		if !matched {
			fmt.Println("URL is not valid")
			return
		}

		html := utils.GetHTML(url)

		links := utils.GetLinks(url, html)

		waitGroup.Add(len(links))

		for _, link := range links {
			go func(link string) {
				defer waitGroup.Done()
				html := utils.GetHTML(link)
				if html == "" {
					return
				}
				utils.SaveFile(link, html, output)
			}(link)
		}
		waitGroup.Wait()
	},
}

func init() {
	RootCmd.AddCommand(mirror)
}
