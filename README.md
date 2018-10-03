## api.moov.io

[![Build Status](https://travis-ci.com/moov-io/api.svg?branch=master)](https://travis-ci.com/moov-io/api)
[![Apache 2 licensed](https://img.shields.io/badge/license-Apache2-blue.svg)](https://raw.githubusercontent.com/moov-io/api/master/LICENSE)

This repository holds our [`api.moov.io`](https://api.moov.io) website, which is our API documentation.

Documentation is markdown generated by [ReDoc](https://github.com/Rebilly/ReDoc). View source on the [Swagger Editor](https://editor.swagger.io/?url=https://raw.githubusercontent.com/moov-io/api/master/openapi.yaml).

## Getting Started / Install

1. Clone the repository somewhere (`git clone git@github.com:moov-io/api.git`)
1. Edit the files to make your change (i.e. `openpai.yaml`)
1. Run `make build` to generate a docker image with your changes
   1. Verify your change looks ok by running the docker image
      - `docker run -p 8080:8080 -it moov/api.moov.io:latest`
      - Load [localhost:8080](http://localhost:8080) in a web browser
1. Commit your changes, push up a new branch, and create a Pull Request!

## Getting Help

 channel | info
 ------- | -------
 Google Group [moov-users](https://groups.google.com/forum/#!forum/moov-users)| The Moov users Google group is for contributors other people contributing to the Moov project. You can join them without a google account by sending an email to [moov-users+subscribe@googlegroups.com](mailto:moov-users+subscribe@googlegroups.com). After receiving the join-request message, you can simply reply to that to confirm the subscription.
Twitter [@moov_io](https://twitter.com/moov_io)	| You can follow Moov.IO's Twitter feed to get updates on our project(s). You can also tweet us questions or just share blogs or stories.
[GitHub Issue](https://github.com/moov-io) | If you are able to reproduce an problem please open a GitHub Issue under the specific project that caused the error.
[moov-io slack](http://moov-io.slack.com/) | Join our slack channel to have an interactive discussion about the development of the project. [Request an invite to the slack channel](https://join.slack.com/t/moov-io/shared_invite/enQtNDE5NzIwNTYxODEwLTRkYTcyZDI5ZTlkZWRjMzlhMWVhMGZlOTZiOTk4MmM3MmRhZDY4OTJiMDVjOTE2MGEyNWYzYzY1MGMyMThiZjg)

## Contributing

Yes please! Please start by reviewing our [Code of Conduct](https://github.com/moov-io/ach/blob/master/CODE_OF_CONDUCT.md).

You only have a fresh set of eyes once! The easiest way to contribute is to give feedback on the documentation that you are reading right now. This can be as simple as sending a message to our Google Group with your feedback or updating the markdown in this documentation and issuing a pull request.

- [moov.io](https://moov.io/)
- [api.moov.io](https://api.moov.io/) (This project)
- [docs.moov.io](https://docs.moov.io/)

## License

Apache License 2.0 See [LICENSE](LICENSE) for details.
