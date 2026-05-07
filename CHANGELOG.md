# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Fixed

- README typo ("sipmle and effecitve" → "simple and effective").

### Changed

- Lower `go.mod` to `go 1.24` (matches the `tool` directive's minimum).
- Drop the "Go 1.24+" claim from the README — `go.mod` is the single
  source of truth for the minimum toolchain.

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
