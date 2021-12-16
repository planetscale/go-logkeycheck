# go-logkeycheck

Lint Go programs that use the `uber/zap` or `planetscale/log` loggers to ensure log field names are consistent.

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

## Usage

Run the linter:

```console
go-logkeycheck ./...
```

## Example

TODO use fatih's faillint as example for example

TODO: ignore example (also test this)