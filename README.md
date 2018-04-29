# emersyx_telegram [![Build Status][build-img]][build-url] [![Go Report Card][gorep-img]][gorep-url] [![GoDoc][godep-img]][godep-url]

This is the vanilla Telegram gateway implementation for the emersyx platform. This gateway acts as both peripheral and
receptor simultaneously.

## Usage

Source files in `emtg` provide the implementation of the go plugin. The plugin can be built by running `make`. The
resulting `emtg.so` file can then be used by emersyx core and router implementations from the [main emersyx
repository][emersyx-repo].

[build-img]: https://travis-ci.org/emersyx/emersyx_telegram.svg?branch=master
[build-url]: https://travis-ci.org/emersyx/emersyx_telegram
[gorep-img]: https://goreportcard.com/badge/github.com/emersyx/emersyx_telegram
[gorep-url]: https://goreportcard.com/report/github.com/emersyx/emersyx_telegram
[godep-img]: https://godoc.org/emersyx.net/emersyx_telegram?status.svg
[godep-url]: https://godoc.org/emersyx.net/emersyx_telegram
[emersyx-repo]: https://github.com/emersyx/emersyx
