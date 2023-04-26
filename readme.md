# Mirroring

A command line tool to mirror a website locally.

![code-snippet](https://user-images.githubusercontent.com/33763843/234628553-00f0456b-8fd8-422c-96bd-654306a1175a.png)


Requires go 1.19 or higher.

### Install

```bash
go get github.com/falcucci/mirroring
```

#### mirror

mirroging mirror – Makes (among other things) the download recursive.

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
 provide "happy-path" test coverage () (missing)
- be implemented in Go;

