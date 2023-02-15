# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased ([diff][unreleased-diff])

## [0.1.0][] ([diff][0.1.0-diff]) - 2023-02-15

### Added

- Added database-related generated methods.
- Added `CHANGELOG.md`, `CODE_OF_CONDUCT.md`, and `CONTRIBUTING.md`.
- Added copyright headers to each file.

### Changed

- Changed the underlying templates from `.mustache` to `.handlebars` format.

### Fixed

- Fixed the panic when setting `VAULT_RATE_LIMIT` environment variable.

## [0.1.0-beta][] ([diff][0.1.0-beta-diff]) - 2023-02-01

### Added

- First public release of the library.

<!-- diffs -->

[unreleased-diff]:
  https://github.com/hashicorp/vault-client-go/compare/v0.1.0...HEAD
[0.1.0-diff]:
  https://github.com/hashicorp/vault-client-go/compare/v0.1.0-beta...v0.1.0
[0.1.0-beta-diff]:
  https://github.com/hashicorp/vault-client-go/commits/v0.1.0-beta

<!-- releases -->

[0.1.0]:
  https://github.com/hashicorp/vault-client-go/releases/tag/v0.1.0
[0.1.0-beta]:
  https://github.com/hashicorp/vault-client-go/releases/tag/v0.1.0-beta
