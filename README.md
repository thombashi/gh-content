# gh-content

[gh][] extension to fetch a content from a GitHub repository.

[![Release](https://img.shields.io/github/release/thombashi/gh-content.svg?style=flat-square)](https://github.com/thombashi/gh-content/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/thombashi/gh-content.svg)](https://pkg.go.dev/github.com/thombashi/gh-content)
[![Go Report Card](https://goreportcard.com/badge/github.com/thombashi/gh-content)](https://goreportcard.com/report/github.com/thombashi/gh-content)
[![CI](https://github.com/thombashi/gh-content/actions/workflows/ci.yaml/badge.svg)](https://github.com/thombashi/gh-content/actions/workflows/ci.yaml)
[![CodeQL](https://github.com/thombashi/gh-content/actions/workflows/github-code-scanning/codeql/badge.svg)](https://github.com/thombashi/gh-content/actions/workflows/github-code-scanning/codeql)


## Installation

```console
gh extension install thombashi/gh-content
```

## Upgrade

```console
gh extension upgrade content
```


## Usage

### Command help

```
      --log-level string   log level (debug, info, warn, error) (default "info")
  -o, --output string      output file path. If not specified, output to stdout.
  -R, --repo string        GitHub repository ID. If not specified, use the current repository.
```

### Examples

```console
gh content --repo thombashi/gh-content README.md
```


[gh]: https://docs.github.com/en/github-cli/github-cli/about-github-cli
