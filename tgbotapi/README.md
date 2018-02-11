# Low level Telegram Bot API bindings for go

This package contains low level bindings for the Telegram Bot API for the go programming language. The bindings follow
as close as possible the names of functions, parameters and return values from the [official Telegram documentation][1].
The purpose of this package is to provide the low level functionality which can further be included into a higher level
library.

This package does not make use of any non-standard third party libraries. It relies solely on types from the standard go
library.

The package consists of the following files:
* `tgbotapi.go` - contains the implementation of the [Bot API methods][2]
* `tgbotapi_test.go` - contains tests which connect to the Telegram back-end, make requests and print the result
* `util.go` - any general purpose utility code

The functions do not format the responses into go types, but instead return raw strings. The strings are meant to be
further processed at upper layers.

The bindings are not complete, but required features should be easy to implement by following the provided examples.

## Obtaining the chat_id value

There are two methods of obtaining the `chat_id` value for a particular group, supergroup or private chat. The first one
involves using the web-based Telegram client. The set of steps is described [here][3]. The second method is to add a bot
to the group, send a message to the group and then follow this API URL:

```
https://api.telegram.org/bot<yourtoken>/getUpdates
```

You will receive JSON formatted data with the correct `chat_id` value.

[1]: https://core.telegram.org/bots/api
[2]: https://core.telegram.org/bots/api#available-methods
[3]: https://github.com/GabrielRF/telegram-id#web-channel-id
