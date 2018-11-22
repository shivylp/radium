# radium

[![GoDoc](https://godoc.org/github.com/spy16/radium?status.svg)](https://godoc.org/github.com/spy16/radium) [![Build Status](https://travis-ci.org/spy16/radium.svg?branch=master)](https://travis-ci.org/spy16/radium) [![Go Report Card](https://goreportcard.com/badge/github.com/spy16/radium)](https://goreportcard.com/report/github.com/spy16/radium)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fspy16%2Fradium.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fspy16%2Fradium?ref=badge_shield)

`radium` is a platform (client and optional server) for viewing
reference articles, cheat sheets etc. right from a shell. `radium`
is written in `Go` (`Golang`)

## Install

1. Download the latest pre-built binary for your platform from [releases](https://github.com/spy16/radium/releases) page.
2. Run the binary! (Optionally, create a `radium.yaml` file in `~` to customize)

## Build

`radium` requires Go 1.8+ to build. Simply run `make` command in
the source directory to build and install the binary into your `$GOPATH/bin`
directory. If you just want to build the binary, run `make build`

## Run
You can run `radium --help` to see the list of available commands.

### Querying from command-line

```bash
radium query "append file in go"

radium query dir --attr platform:windows

radium query go
```


> `--attr` is not part of radium framework but part of the source
> implementation itself. Weather to consume the attributes or not
> is decided by the source implementation.

### Querying from curl

For this, you need to run `radium` in server mode first using the
command: `radium serve --addr=localhost:8080`

Then

```bash
curl "localhost:8080/search?q=append+file+in+go"

curl "localhost:8080/search?q=dir&platform=windows"

curl "localhost:8080/search?q=go"
```

> When using http api, all query parameters except `q` and `strategy` will be
> assumed to be attributes

### Running as Clipboard Monitor

Run `radium serve --clipboard` to start radium in server+clipboard
mode (pass `--addr=""` to run in clipboard-only mode).

Now, everytime you copy some text into clipboard (which is less than
5 words), `radium` is going to run a query and try to find some results.
If a result is found, it will be pasted back into the clipboard

## How it works?

`radium` works by querying/scraping different knowledge sources
(e.g. tldr-pages, LearnXInYMinutes etc.). A `Source` in `radium`
is a `Go` interface and can be implemented to add new references
to provide more relevant results.

```go
type Source interface {
  Search(query Query) ([]Article, error)
}
```

Currently following implementations are available:


1. `sources.TLDR` using the awesome [tldr-pages](https://github.com/tldr-pages/tldr) project
2. `sources.LearnXInY` using the awesome [Learn X In Y Minutes](https://github.com/adambard/learnxinyminutes-docs) project
3. `cheatsh.CheatSh` using the awesome [cheat.sh](https://github.com/chubin/cheat.sh) project
4. `sources.Radium` which can be used to query other `radium` servers to enable distributed setup
5. `wikipedia.Wikipedia` which queries and extracts first paragraph from [Wikipedia](https://en.wikipedia.org)
6. `duckduckgo.DuckDuckGo` which uses [DuckDuckGo Instant Answer API](https://api.duckduckgo.com/)

See [sources/README.md](./sources/README.md) for more information

## TODO:

- [x] Make sources configurable
  - [x] sources to be used should be configurable per instance
- [ ] a configurable caching mechanism to enable offline usage
- [ ] Add more sources
  - [x] Wikipedia
  - [x] DuckDuckGo
- [ ] Enable markdown to console colored output ?
- [ ] Enable clipboard monitoring
  - [x] everytime user copies a string, run radium query
  - [x] if a result is available within certain time window, replace the clipboard
    content with the solution
  - [ ] enable query only if clipboard text is in special format to reduce unwanted paste-backs


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fspy16%2Fradium.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fspy16%2Fradium?ref=badge_large)