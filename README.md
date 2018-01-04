# emersyx_telegram [![Build Status][build-img]][build-url] [![Go Report Card][gorep-img]][gorep-url]

Telegram receptor and resource for emersyx.

## Build

Source files in `emtg_impl` provide the implementation of the go plugin. They have to be built using the command:

```
go build -buildmode=plugin -o emtg.so emtg/*
```

The resulting `emtg.so` file can then be used by emersyx core.

## Notes

The `TelegramBot` struct follows the APIs defined in the [emersyx_apis][1] repository, specifically those from the
[emtgapi][2] folder.

The `NewTelegramBot` function must be used to create new `TelegramBot` instances. An example of how to use this function
can be found in the `emtg/tgbot_test.go` file.

## tgbotapi

The core functionality for interacting with the Telegram Bot back-end is provided by the `tgbotapi` package. This
package provides low level access to the back-end using only functionality from the standard go library. This package is
not directly used by emersyx and in theory can be re-used in other projects.

For example usage of the `tgbotapi` package, check the `tgbotapi/tgbotapi_test.go` file.

[build-img]: https://travis-ci.org/emersyx/emersyx_telegram.svg?branch=master
[build-url]: https://travis-ci.org/emersyx/emersyx_telegram
[gorep-img]: https://goreportcard.com/badge/github.com/emersyx/emersyx_telegram
[gorep-url]: https://goreportcard.com/report/github.com/emersyx/emersyx_telegram
[1]: https://github.com/emersyx/emersyx_apis
[2]: https://github.com/emersyx/emersyx_apis/tree/master/emtgapi
