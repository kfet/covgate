# covgate

The smallest possible Go coverage gate: one binary, one regex file, one
threshold flag.

`covgate` reads a Go coverage profile, drops lines matching regexes in a
`.covignore` file, writes the filtered profile back out, and fails if
the resulting statement-weighted coverage is below a minimum threshold.

## Install

As a Go tool (recommended):

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
syntax). The pattern is matched against the raw profile line, which
looks like:

```
github.com/you/pkg/file.go:12.34,15.6 7 0
```

So `^github.com/you/pkg/file\.go:` excludes the whole file;
`/generated\.go:` excludes any file named `generated.go` in any
package; `^github.com/you/proj/e2e/` excludes a whole subpackage;
and so on.

Blank lines and `#` comments are skipped. Example:

```
# Entry-point shims — bare flag parsing & dependency wiring.
^github.com/you/proj/cmd/[^/]+/main\.go:

# Generated code.
/zz_generated\.go:

# Integration-only test package, exercised by a separate suite.
^github.com/you/proj/e2e/
```

### One rule: file- or directory-level patterns only

The single discipline `covgate` is opinionated about is: **express
exclusions at the file or directory boundary — never line numbers,
never per-function regexes**. Both rot the moment surrounding code
shifts, and per-function regexes silently mask new untested code
added inside the same function.

Anything coarser is fine. Whole files, whole packages, whole
subdirectories, whole generated bundles — all legitimate. There is
no "approved list" of filenames; pick patterns that match how
*your* code is organised.

A genuinely useful boundary is a **thin wrapper around an unmockable
dependency**, isolated in its own file or directory (commonly named
`ext`, `external`, `bridge`, etc.). The wrapper exists to keep the
unmockable surface narrow so the rest of the package can be tested
through it; the wrapper itself is exercised only by integration
tests. Exclude that file or directory wholesale and move on:

```
# Thin wrapper around an unmockable C lib / OS API / vendor SDK.
^github.com/you/proj/internal/foo/ext\.go:
^github.com/you/proj/internal/bar/ext/
```

#### Anti-pattern to avoid: a dedicated "unreachable.go" file

A common temptation is to invent a per-package `unreachable.go` (or
`untestable.go`) file just to give the gate something to exclude.
Don't. That's the same shape as a `utils.go` grab-bag: code
organised by a meta-property ("hard to test") rather than by what
it *is*. The `ext.go` / `ext/` pattern above is different — it
isolates a *real domain boundary* (the unmockable dependency); a
file named after the testing problem itself is just a dumping
ground.

When tempted to add an exclusion, prefer restructuring first:

- Push impossible-error branches into a `must…` helper that
  panics, so callers have nothing to cover.
- Move the genuinely-isolated dependency behind a thin wrapper at
  a real domain boundary (`ext.go`, `osfs.go`, `vendor/foo/`),
  and exclude the wrapper.

Only fall back to a file/dir exclusion when the boundary is real.

## Why regex on profile lines?

It's simple and effective. Other tools in this space either:

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

MIT
