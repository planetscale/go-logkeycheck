# go-logkeycheck

Lint Go programs that use the `go.uber.org/zap` or `planetscale/log` loggers to ensure log field names are consistent.

Current linting rules:

1. Log field names should be `snake_case`.

## Install

Install latest:

```console
go install github.com/planetscale/go-logkeycheck@latest
```

Or specify a version:

```console
go install github.com/planetscale/go-logkeycheck@v0.0.1
```

Refer to [github releases](https://github.com/planetscale/go-logkeycheck/releases) for a list of available versions.

## Example

Assume we have the following file:

```go
package main

import (
  "github.com/planetscale/log"
  // "go.uber.org/zap" is also supported
)

func main() {
  logger := log.NewPlanetScaleLogger()
  defer logger.Sync()

  logger.Info("info log with fields",
    log.String("user_id", "12345678"),
    log.String("branchId", "abcdefghijkl"),
  )
```

Run `go-logkeycheck` to check if our log fields are in the correct format:

```console
$ go-logkeycheck ./...
main.go:13:17: log key 'branchId' should be snake_case.
```

## Usage

Run the linter on a file, directory, or Go package:

```console
go-logkeycheck foo.go # lint a single file
go-logkeycheck ./...  # recursively lint all files
```

## golangci-lint plugin support

> Support for running in plugin mode is experimental.

The linter can be compiled as a plugin for `golangci-lint`.

You will need to compile a `CGO_ENABLED=1` version of `golangci-lint`. The official binaries are not compatible with plugins.

Run `make build-plugin` to generate the plugin object `go-logkeycheck.so`.

Include the following in your `golangci.yml` file:

```yaml
linters-settings:
  custom:
    logkeycheck:
      path: ./go-logkeycheck.so

linters:
  enable:
    - logkeycheck
```

## Credits

This tool is built on top of the `go/analysis` package that makes it easy to write customer linters in Go.

If you are interested in writing a linter for Go check out these excellent resources which were invaluable in writing this tool:

- @fatih's [Using go/analysis to write a custom linter](https://arslan.io/2019/06/13/using-go-analysis-to-write-a-custom-linter/)
- [Custom linter plugins for golangci-lint](https://tech.devoted.com/custom--linter-plugins-for--golangci-lint-cf56b9091491)
