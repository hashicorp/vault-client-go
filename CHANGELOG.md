# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased ([diff][unreleased-diff])

## [0.3.1][] ([diff][0.3.1-diff]) - 2023-05-10

### Changed

- Improved documentation for request logging in README.

### Fixed

- Fixed bug #147 -- the empty generated fields are now properly omitted.
- Fixed mount path examples in the generated documentation.
- Fixed the Makefile to properly clean up before regeneration.

## [0.3.0][] ([diff][0.3.0-diff]) - 2023-05-04

### Added

- Added response structures for `sys`, `kv-v1`, `kv-v2`, `pki`, and `approle`.
- Added support for non-string query parameters.
- Added prettier README formatter to the Makefile and GitHub actions.
- Added security note to README.

### Removed

- Removed redundant methods (e.g. `TokenLookupSelf2`, `TokenLookupSelf3`, etc).
- Removed endpoints for `ad` (consolidated into `ldap`).
- Removed endpoints for `openldap` (alias of `ldap`).
- Removed endpoints for `oidc` (consolidated into `jwt`).
- Removed endpoints for `pfc` (deprecated plugin).

### Changed

- Regenerated with new method names, request names, and response names.

### Fixed

- Fixed mount path logic in method signatures.

## [0.2.0][] ([diff][0.2.0-diff]) - 2023-03-01

### Changed

- Renamed `Address` and `HTTPClient`.

### Fixed

- The mount path logic is now reliably generated using `.handlebars` templates.
- Fixed an issue with how model documents were generated.

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
  https://github.com/hashicorp/vault-client-go/compare/v0.3.1...HEAD
[0.3.1-diff]:
  https://github.com/hashicorp/vault-client-go/compare/v0.3.0...v0.3.1
[0.3.0-diff]:
  https://github.com/hashicorp/vault-client-go/compare/v0.2.0...v0.3.0
[0.2.0-diff]:
  https://github.com/hashicorp/vault-client-go/compare/v0.1.0...v0.2.0
[0.1.0-diff]:
  https://github.com/hashicorp/vault-client-go/compare/v0.1.0-beta...v0.1.0
[0.1.0-beta-diff]:
  https://github.com/hashicorp/vault-client-go/commits/v0.1.0-beta

<!-- releases -->

[0.3.1]: https://github.com/hashicorp/vault-client-go/releases/tag/v0.3.1
[0.3.0]: https://github.com/hashicorp/vault-client-go/releases/tag/v0.3.0
[0.2.0]: https://github.com/hashicorp/vault-client-go/releases/tag/v0.2.0
[0.1.0]: https://github.com/hashicorp/vault-client-go/releases/tag/v0.1.0
[0.1.0-beta]:
  https://github.com/hashicorp/vault-client-go/releases/tag/v0.1.0-beta
