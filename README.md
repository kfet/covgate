# covgate

The smallest possible Go coverage gate: one binary, one regex file, one
threshold flag.

`covgate` reads a Go coverage profile, drops lines matching regexes in a
`.covignore` file, writes the filtered profile back out, and fails if
the resulting statement-weighted coverage is below a minimum threshold.

That's it. No YAML config, no GitHub Action wrapper, no badges, no
per-package thresholds. If you want any of those, use
[vladopajic/go-test-coverage](https://github.com/vladopajic/go-test-coverage).

## Install

As a Go tool (Go 1.24+, recommended):

```sh
go get -tool github.com/kfet/covgate/cmd/covgate@latest
```

Then invoke as `go tool covgate …`.

Or as a regular binary:

```sh
go install github.com/kfet/covgate/cmd/covgate@latest
```

## Use

```sh
go test -covermode=set -coverprofile=coverage.tmp.out ./...
go tool covgate \
    -profile=coverage.tmp.out \
    -out=coverage.out         \
    -ignore=.covignore        \
    -min=100
```

On success: prints `coverage: 100.0% of statements (N/N)` to stdout,
exits 0, and writes the filtered profile to `coverage.out` for use
with `go tool cover -html=coverage.out` etc.

On failure: prints the gap, lists each surviving uncovered profile
entry, and exits 1.

## `.covignore` format

One regular expression per line (Go's [`regexp`](https://pkg.go.dev/regexp)
syntax). The pattern is matched against the **raw profile line**, which
looks like:

```
github.com/you/pkg/file.go:12.34,15.6 7 0
```

So `^github.com/you/pkg/file\.go:` excludes the whole file;
`/generated\.go:` excludes any file named `generated.go` in any
package; and so on.

Blank lines and `#` comments are skipped. Example:

```
# Entry-point shims — bare flag parsing & dependency wiring.
^github.com/you/proj/cmd/[^/]+/main\.go:

# Structurally-unreachable defensive code, isolated per package.
/unreachable\.go:
```

The recommended discipline is to express exclusions as **file-level
patterns only** — never line numbers or per-function regexes, both of
which silently rot when surrounding code changes. Move unreachable code
into a dedicated file (e.g. `unreachable.go`) and exclude that file.

## Why regex on profile lines?

Other tools in this space either:

- match globs against package paths (PaloAltoNetworks/cov), or
- require a YAML config and cover the full CI lifecycle (vladopajic/go-test-coverage), or
- use source-comment annotations like `//coverage:ignore` (hexira/go-ignore-cov).

`covgate` matches regexes against the literal profile lines that `go
test -coverprofile` produces. That's the lowest-level surface possible:
you can express any exclusion `grep -v -E` could, and there's no
metadata layer to fall out of sync with the code.

## As a library

```go
import "github.com/kfet/covgate"

err := covgate.Run(covgate.Config{
    ProfilePath: "coverage.tmp.out",
    OutPath:     "coverage.out",
    IgnorePath:  ".covignore",
    Min:         100,
    Stdout:      os.Stdout,
    Stderr:      os.Stderr,
})
```

## Makefile recipe

```makefile
coverage:
	go test -covermode=set -coverprofile=bin/coverage.tmp.out ./...
	go tool covgate -profile=bin/coverage.tmp.out \
	                -out=bin/coverage.out         \
	                -ignore=.covignore            \
	                -min=100
```

## License

MIT.
