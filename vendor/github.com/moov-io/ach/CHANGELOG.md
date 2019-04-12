## v1.0.0 (Released 2019-03-26)

- No changes from v0.6.0

## v0.6.0 (Released 2019-03-26)

BREAKING CHANGES

- `file.NotificationOfChange` accepts `Batcher` instead of `*BatchCOR` to comply with linter errors

ADDITIONS

- Add const values for `BatchHeader.StandardEntryClassCode` (See [#392](https://github.com/moov-io/ach/issues/392))
- Add const values for `BatchHeader.ServiceClassCode` and `BatchControl.ServiceClassCode`. (See [#391](https://github.com/moov-io/ach/issues/391))
- Add const values for `EntryDetail.TransactionCode` (See [#363](https://github.com/moov-io/ach/issues/363))
- server: Record `ach_files_deleted` metric. (See: [#408](https://github.com/moov-io/ach/pull/408))
- server: log x-request-id header if present. (See: [#407](https://github.com/moov-io/ach/pull/407))
- server: Delete old `ach.File` objects from in-memory repository according to `ACH_FILE_TTL` env variable.
- server: Support `-log.format=json` for JSON formatted logs

BUG FIXES

- Accept File's without an ID specified. Generate a random ID. (See: [#405](https://github.com/moov-io/ach/pull/405))
- server: Fix nil panics. (See: [#406](https://github.com/moov-io/ach/pull/406))
- server: Fix type-casting panics. (See: [#423](https://github.com/moov-io/ach/pull/423))
- server: validate file endpoint returns 400 instead of 500 (See: [#488](https://github.com/moov-io/ach/pull/488))
- server: set CORS headers on `GET /ping` route

BUILD

- `megacheck` is deprecated. staticcheck should be used instead. (See [#430](https://github.com/moov-io/ach/issues/430))
- Automate releases with Docker and binary uploads to release page.
- Update dependencies to their latest versions
- Update to Go 1.12

## v0.5.0 (Released 2018-11-29)

BREAKING CHANGES

- `TraceNumber` has been changed from `int` to a `string`. (See [#366](https://github.com/moov-io/ach/issues/366))
   - Previously zero-prefixed ABA routing numbers would have their leading zero truncated.
- `OriginalTrace` has been changed from `int` to a `string`. (See [#366](https://github.com/moov-io/ach/issues/366))

ADDITIONS

- Support `StandardEntryClassCode` (Batch types):
    - ADV (See [#340](https://github.com/moov-io/ach/issues/340))
    - TRC (See [#346](https://github.com/moov-io/ach/issues/346))
    - TRX (See [#372](https://github.com/moov-io/ach/issues/372))
    - XCK (See [#347](https://github.com/moov-io/ach/issues/347))
- `TransactionCode` match `ServiceClassCode` (See [#56](https://github.com/moov-io/ach/issues/56))
- `Addenda02.TerminalState` validation for BatchPOS and BatchSHR (See [#375](https://github.com/moov-io/ach/issues/375))

REMOVALS

- Remove deprecated functions from `EntryDetail` (See [#385](https://github.com/moov-io/ach/issues/385))

## v0.4.0 (Released 2018-11-06)

BREAKING CHANGES

- `EntryDetail.Addendum` has been broken out into `Addenda02`, `Addenda05`, `Addenda98`, and `Addenda99` fields on `EntryDetail`.
- IAT `EntryDetail.Addendum` has been broken out into Addenda 10-18, 98 and 99.

ADDITIONS

- Support `StandardEntryClassCode` (Batch types):
  - ACK (See [#327](https://github.com/moov-io/ach/issues/327))
  - ATX (See [#327](https://github.com/moov-io/ach/issues/327))
  - DNE (See [#342](https://github.com/moov-io/ach/issues/342))
  - ENR (See [#343](https://github.com/moov-io/ach/issues/343))
- Support NOC for IAT Entries (See [#328](https://github.com/moov-io/ach/issues/328))
- Add `FileFromJson` for reading `File` objects as JSON.
- Add `X-Total-Count` response headers on `GET /files/:id/batches` (See [#280](https://github.com/moov-io/ach/issues/280))

IMPROVEMENTS

- Multiple parsing errors are returned in `Reader.Read() error`'s error message.
- IAT: Validate ISODestinationCountryCode according to ISO 3166
- IAT: Validate ISOOriginatingCurrencyCode and ISODestinationCurrencyCode according to ISO 4217
- build: Run CI tests against Windows
- Verify record lengths in [IAT] BatchHeader, [IAT] BatchControl, FileControl, FileHeader, and [IAT] EntryDetail.

BUG FIXES
- `cmd/server`: don't expect trailing slash on endpoints
- `cmd/server`: Grab write lock on delete requests
- Several panics are fixed from fuzzing
- Reject invalid ABA routing numbers

## v0.3.3 (Released 2018-10-08)

ADDITIONS

- Expose `TypeCode` on Addenda records

IMPROVEMENTS

- Run as unprivileged user in Docker image
- `cmd/readACH`: add -json to print `File` as JSON
- `cmd/writeACH`: validate file before writing
- `cmd/writeACH`: add -json to output `File` as JSON

BUG FIXES

- Fix reading batches out from JSON endpoint
- Fix plaintext ACH file rendering endpoint

## v0.3.2 (Released 2018-10-05)

ADDITIONS

- Handle pre-flight requests and proxy back CORS headers

## v0.3.1 (Released 2018-10-05)

ADDITIONS

- Add `X-Total-Count` on GET responses
- Proxy back CORS headers if sent on the request

BUG FIXES

- Drop requirement for /-suffix on GET /files
- Don't trample content-type in encodeResponse

## v0.3.0 (Released 2018-09-26)

FEATURES

- Added HTTP Server
- SEC Code CIE (Customer-Initiated Entry) ([#209](https://github.com/moov-io/ach/issues/209))
- Support IAT ([#211](https://github.com/moov-io/ach/issues/211))
- IAT Returns ([#233](https://github.com/moov-io/ach/issues/233))
- Support CTX ([#212](https://github.com/moov-io/ach/issues/212))

IMPROVEMENTS

- Added admin HTTP service.
- Added `GET /ping` route (unauthed)

BUG FIXES

- Fixes to parsing and validation.
