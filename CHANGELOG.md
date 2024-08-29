# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.0.5] - 2024-08-29

### Removed

- Remove `PolicyExceptionDraft` API and CRD.

## [0.0.4] - 2024-08-13

### Changed

- Change `PolicyConfig` shortNames to avoid collision with `Policy` CRD.

## [0.0.3] - 2024-08-07

### Added

- Add API definitions for `Policy` and `PolicyConfig`.

### Modified

- Add shortNames and modify required fields for `PolicyManifests` and `AutomatedExceptions`.

## [0.0.2] - 2024-08-06

### Changed

- Make PolicyManifest fields `Exception` and `AutomatedExceptions` optional.

## [0.0.1] - 2024-07-30

### Added

- First release of the `policy-api` project.

[Unreleased]: https://github.com/giantswarm/policy-api/compare/v0.0.5...HEAD
[0.0.5]: https://github.com/giantswarm/policy-api/compare/v0.0.4...v0.0.5
[0.0.4]: https://github.com/giantswarm/policy-api/compare/v0.0.3...v0.0.4
[0.0.3]: https://github.com/giantswarm/policy-api/compare/v0.0.2...v0.0.3
[0.0.2]: https://github.com/giantswarm/policy-api/compare/v0.0.1...v0.0.2
[0.0.1]: https://github.com/giantswarm/policy-api/releases/tag/v0.0.1
