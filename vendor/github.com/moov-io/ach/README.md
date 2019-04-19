moov-io/ach
===
[![GoDoc](https://godoc.org/github.com/moov-io/ach?status.svg)](https://godoc.org/github.com/moov-io/ach)
[![Build Status](https://travis-ci.com/moov-io/ach.svg?branch=master)](https://travis-ci.com/moov-io/ach)
[![Coverage Status](https://codecov.io/gh/moov-io/ach/branch/master/graph/badge.svg)](https://codecov.io/gh/moov-io/ach)
[![Go Report Card](https://goreportcard.com/badge/github.com/moov-io/ach)](https://goreportcard.com/report/github.com/moov-io/ach)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/ach/master/LICENSE)

Package `github.com/moov-io/ach` implements a file reader and writer written in Go along with a HTTP API for creating, parsing and validating Automated Clearing House ([ACH](https://en.wikipedia.org/wiki/Automated_Clearing_House)) files. ACH is the primary method of electronic money movement throughout the United States.

Docs: [docs.moov.io](https://docs.moov.io/en/latest/) | [api docs](https://api.moov.io)

## Project Status

Moov ACH is under active development and in production for multiple companies. Please star the project if you are interested in its progress. We've built an HTTP API for ACH file creation and validation. Currently we support generating and parsing all Standard Entry Class (SEC) codes. If you have layers above ACH to simplify tasks or found bugs we would appreciate an issue or pull request. Thanks!

## Usage

### Go library

`github.com/moov-io/ach` offers a Go based ACH file reader and writer. To get started checkout a specific example:

<details>
<summary>Supported Standard Entry Class (SEC) codes</summary>

| SEC Code | Name                                  | Read Example                      | Write Example                      |
|----------|---------------------------------------|-----------------------------------|------------------------------------|
| ACK      | Acknowledgment Entry for CCD          | [Link](test/ach-ack-read/main.go) | [Link](test/ach-ack-write/main.go) |
| ADV      | Automated Accounting Advice           | [Link](test/ach-adv-read/main.go) | [Link](test/ach-adv-write/main.go) |
| ARC      | Accounts Receivable Entry             | [Link](test/ach-arc-read/main.go) | [Link](test/ach-arc-write/main.go) |
| ATX      | Acknowledgment Entry for CTX          | [Link](test/ach-atx-read/main.go) | [Link](test/ach-atx-write/main.go) |
| BOC      | Back Office Conversion                | [Link](test/ach-boc-read/main.go) | [Link](test/ach-boc-write/main.go) |
| CCD      | Corporate credit or debit             | [Link](test/ach-ccd-read/main.go) | [Link](test/ach-ccd-write/main.go) |
| CIE      | Customer-Initiated Entry              | [Link](test/ach-cie-read/main.go) | [Link](test/ach-cie-write/main.go) |
| COR      | Automated Notification of Change(NOC) | [Link](test/ach-cor-read/main.go) | [Link](test/ach-cor-write/main.go) |
| CTX      | Corporate Trade Exchange              | [Link](test/ach-ctx-read/main.go) | [Link](test/ach-ctx-write/main.go) |
| DNE      | Death Notification Entry              | [Link](test/ach-dne-read/main.go) | [Link](test/ach-dne-write/main.go) |
| ENR      | Automatic Enrollment Entry            | [Link](test/ach-enr-read/main.go) | [Link](test/ach-enr-write/main.go) |
| IAT      | International ACH Transactions        | [Link](test/ach-iat-read/main.go) | [Link](test/ach-iat-write/main.go) |
| MTE      | Machine Transfer Entry                | [Link](test/ach-mte-read/main.go) | [Link](test/ach-mte-write/main.go) |
| POP      | Point of Purchase                     | [Link](test/ach-pop-read/main.go) | [Link](test/ach-pop-write/main.go) |
| POS      | Point of Sale                         | [Link](test/ach-pos-read/main.go) | [Link](test/ach-pos-write/main.go) |
| PPD      | Prearranged payment and deposits      | [Link](test/ach-ppd-read/main.go) | [Link](test/ach-ppd-write/main.go) |
| RCK      | Represented Check Entries             | [Link](test/ach-rck-read/main.go) | [Link](test/ach-rck-write/main.go) |
| SHR      | Shared Network Entry                  | [Link](test/ach-shr-read/main.go) | [Link](test/ach-shr-write/main.go) |
| TRC      | Check Truncation Entry                | [Link](test/ach-trc-read/main.go) | [Link](test/ach-trc-write/main.go) |
| TRX      | Check Truncation Entries Exchange     | [Link](test/ach-trx-read/main.go) | [Link](test/ach-trx-write/main.go) |
| TEL      | Telephone-Initiated Entry             | [Link](test/ach-tel-read/main.go) | [Link](test/ach-tel-write/main.go) |
| TRC      | Truncated Check Entry                 | [Link](test/ach-trc-read/main.go) | [Link](test/ach-trc-write/main.go) |
| WEB      | Internet-initiated Entries            | [Link](test/ach-web-read/main.go) | [Link](test/ach-web-write/main.go) |
| XCK      | Destroyed Check Entry                 | [Link](test/ach-xck-read/main.go) | [Link](test/ach-xck-write/main.go) |

</details>

### HTTP API

`github.com/moov-io/ach/server` offers a HTTP and JSON API for creating and editing files. If you're using Go the `ach.File` type can be used, otherwise just send properly formatted JSON. We have an [example JSON file](test/testdata/ppd-valid.json), but each SEC type will generate differnet JSON.

Examples: [Go](examples/http/main.go) | [Ruby](https://github.com/moov-io/ruby-ach-demo)

- [Create an ACH file for a payment and get the raw file](https://github.com/moov-io/ruby-ach-demo)

## Getting Started

- [Running ACH Server](https://docs.moov.io/en/latest/tutorials/ach-server/)
- [ACH Server metrics](documentation/metrics.md)

### From Source

This project uses [Go Modules](https://github.com/golang/go/wiki/Modules) and thus requires Go 1.11+. You can download the source code and we offer [tagged and released versions](https://github.com/moov-io/ach/releases) as well. We highly recommend you use a tagged release for production.

```
$ git@github.com:moov-io/ach.git

# Pull down into the Go Module cache
$ go get -u github.com/moov-io/ach

$ go doc github.com/moov-io/ach BatchHeader
```

### Configuration

| Environmental Variable | Description | Default |
|-----|-----|-----|
| `ACH_FILE_TTL` | Time to live (TTL) for `*ach.File` objects stored in the in-memory repository. | 0 = No TTL / Never delete files (Example: `240m`) |

## Getting Help

If you have ACH specific questions NACHA (National Automated Clearing House Association) has their [complete specification](documentation/2013-Corporate-Rules-and-Guidelines.pdf) for all file formats and message types.

 channel | info
 ------- | -------
 [Project Documentation](https://docs.moov.io/en/latest/) | Our project documentation available online.
 Google Group [moov-users](https://groups.google.com/forum/#!forum/moov-users)| The Moov users Google group is for contributors other people contributing to the Moov project. You can join them without a google account by sending an email to [moov-users+subscribe@googlegroups.com](mailto:moov-users+subscribe@googlegroups.com). After receiving the join-request message, you can simply reply to that to confirm the subscription.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io) | If you are able to reproduce an problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](http://moov-io.slack.com/) | Join our slack channel to have an interactive discussion about the development of the project. [Request an invite to the slack channel](https://join.slack.com/t/moov-io/shared_invite/enQtNDE5NzIwNTYxODEwLTRkYTcyZDI5ZTlkZWRjMzlhMWVhMGZlOTZiOTk4MmM3MmRhZDY4OTJiMDVjOTE2MGEyNWYzYzY1MGMyMThiZjg)

## Supported and Tested Platforms

- 64-bit Linux (Ubuntu, Debian), macOS, and Windows
- Rasberry Pi

Note: 32-bit platforms have known issues and are not supported.

## Contributing

Yes please! Please review our [Contributing guide](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md) to get started!

Note: This project uses Go Modules, which requires Go 1.11 or higher, but we ship the vendor directory in our repository.

### Releasing

To make a release of ach simply open a pull request with `CHANGELOG.md` and `version.go` updated with the next version number and details. You'll also need to push the tag (i.e. `git push origin v1.0.0`) to origin in order for CI to make the release.

### Fuzzing

We currently run fuzzing over ACH in the form of a [`moov/achfuzz`](https://hub.docker.com/r/moov/achfuzz) Docker image. You can [read more](./test/fuzz-reader/README.md) or run the image and report crasher examples to [`security@moov.io`](mailto:security@moov.io). Thanks!

## License

Apache License 2.0 See [LICENSE](LICENSE) for details.
