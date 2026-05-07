# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

## [0.1.2] - 2026-05-07

### Changed

- Drop the `tool` directive from `go.mod` and lower the minimum Go
  version to 1.21 for wider compatibility. Users on Go 1.24+ can still
  register covgate as a project-local tool via
  `go get -tool github.com/kfet/covgate/cmd/covgate@latest`.

### Fixed

- README typo ("sipmle and effecitve" → "simple and effective").

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
