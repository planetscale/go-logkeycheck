# go-logkeycheck

Lint Go programs that use the `uber/zap` or `planetscale/log` loggers to ensure log field names are consistent.

Current linting rules:

1. Log field names should be `snake_case`.

## Usage

```console
go install ... TODO
```

```console
go-logkeycheck ./...
```