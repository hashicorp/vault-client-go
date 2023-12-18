# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## Unreleased ([diff][unreleased-diff])

## [0.4.3][] ([diff][0.4.3-diff]) - 2023-12-15

### Fixed

- Fixed error parsing for errors with non-standard response bodies (#247).

## [0.4.2][] ([diff][0.4.2-diff]) - 2023-10-19

### Fixed

- Fixed `SetRequestCallbacks` & `SetResponseCallbacks` functions (#242).

## [0.4.1][] ([diff][0.4.1-diff]) - 2023-09-27

### Fixed

- Fixed `SetCustomHeaders` & `WithCustomHeaders` functions (#237).

## [0.4.0][] ([diff][0.4.0-diff]) - 2023-09-18

### Added

- Added a way to pass extra options to `openapi-generator-cli` (#204).
- Added support for custom non-HTTP transports.

### Changed

- A number of generated method names and request/response schemas have changed.
- Renamed `WithCustomQueryParameters` => `WithQueryParameters` (#217).
- Changed how request modifiers behave for slices and maps (#225).
- Switched to Go 1.21 (#234).

### Removed

- Removed support for SRV DNS lookups (#222).
- Removed `ReadWithParameters` (previously deprecated in 0.3.2) (#189).
- Removed `ReadRawWithParameters` (previously deprecated in 0.3.2) (#189).
- Removed `DeleteWithParameters` (previously deprecated in 0.3.2) (#189).

### Fixed

- Improved exclusion of operation IDs (#215).
- Improved generated documentation (#223, #232).
- Improved performance of request modifiers (#224).
- Fixed logic for methods with both `GET` and `LIST` (e.g. `KvV1List`) (#197).
- Fixed warnings during generation due to missing summary and description (#178).
- Fixed small issues in GitHub actions (#182).
- Fixed `TokenCreate*` generated methods (#192).
- Fixed duplication of `sys/raw/` and `sys/leases/lookup/` APIs (#203).
- Fixed templates to deal with arbitrary input APIs (#205).
- Fixed generation of `Query: true` parameters (#206).
- Fixed `CubbyholeWrite` with `TakesArbitraryInput` logic (#206).
- Fixed query parameter casing (#207).
- Fixed query parameters being double-encoded (#213).
- Fixed `/sys/health` to no longer return errors on 4xx & 5xx (#220, #221).
- Fixed `vault.WithRetryConfiguration` to not error on missing parameters.

## [0.3.3][] ([diff][0.3.3-diff]) - 2023-05-26

### Fixed

- Fixed PkiListResponse encoding issue (#175).

## [0.3.2][] ([diff][0.3.2-diff]) - 2023-05-16

### Added

- Added support for custom query parameters.

### Deprecated

- Deprecated `ReadWithParameters`.
- Deprecated `ReadRawWithParameters`.
- Deprecated `DeleteWithParameters`.

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
  https://github.com/hashicorp/vault-client-go/compare/v0.4.3...HEAD
[0.4.3-diff]:
  https://github.com/hashicorp/vault-client-go/compare/v0.4.2...v0.4.3
[0.4.2-diff]:
  https://github.com/hashicorp/vault-client-go/compare/v0.4.1...v0.4.2
[0.4.1-diff]:
  https://github.com/hashicorp/vault-client-go/compare/v0.4.0...v0.4.1
[0.4.0-diff]:
  https://github.com/hashicorp/vault-client-go/compare/v0.3.3...v0.4.0
[0.3.3-diff]:
  https://github.com/hashicorp/vault-client-go/compare/v0.3.2...v0.3.3
[0.3.2-diff]:
  https://github.com/hashicorp/vault-client-go/compare/v0.3.1...v0.3.2
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

[0.4.3]: https://github.com/hashicorp/vault-client-go/releases/tag/v0.4.3
[0.4.2]: https://github.com/hashicorp/vault-client-go/releases/tag/v0.4.2
[0.4.1]: https://github.com/hashicorp/vault-client-go/releases/tag/v0.4.1
[0.4.0]: https://github.com/hashicorp/vault-client-go/releases/tag/v0.4.0
[0.3.3]: https://github.com/hashicorp/vault-client-go/releases/tag/v0.3.3
[0.3.2]: https://github.com/hashicorp/vault-client-go/releases/tag/v0.3.2
[0.3.1]: https://github.com/hashicorp/vault-client-go/releases/tag/v0.3.1
[0.3.0]: https://github.com/hashicorp/vault-client-go/releases/tag/v0.3.0
[0.2.0]: https://github.com/hashicorp/vault-client-go/releases/tag/v0.2.0
[0.1.0]: https://github.com/hashicorp/vault-client-go/releases/tag/v0.1.0
[0.1.0-beta]:
  https://github.com/hashicorp/vault-client-go/releases/tag/v0.1.0-beta
