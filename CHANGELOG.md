# Changelog

All notable changes to this project will be documented in this file.

## [0.1.1] - 2026-05-07

### Changed

- README cleanup.

## [0.1.0] - 2026-05-07

Initial release.

### Added

- `covgate` CLI: filter a Go coverage profile through a `.covignore`
  file (line-oriented regexes, `#` comments) and enforce a minimum
  statement-weighted coverage threshold.
- Library API: `covgate.Run(covgate.Config{…})`.
- Registerable as a Go tool (`tool github.com/kfet/covgate/cmd/covgate`)
  for invocation as `go tool covgate`.
