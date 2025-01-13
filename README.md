# gh-content

[gh][] extension to fetch a content from a GitHub repository.


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
