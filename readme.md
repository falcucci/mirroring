# Mirroring

A command line tool to mirror a website locally.

Requires go 1.19 or higher.

mirroging mirror – Makes (among other things) the download recursive.

#### mirror

Automatically crop away static borders as much as you need
```bash
go run main.go mirror <url> <dir>
```

example:

```
go run main.go mirror https://www.truvity.com ./output
```

How it works?

- accepts a starting URL and a destination directory;
- download the page of the URL fetched;
- save it in the destination directory;
- recursively proceed to any valid links in this URL;
- should correctly handle being interrupted by Ctrl-C;
- perform work in parallel where reasonable;
- support resume functionality by checking the destination directory for downloaded pages;
- skip downloading and processing where not necessary
 provide "happy-path" test coverage ()
- be implemented in Go;

