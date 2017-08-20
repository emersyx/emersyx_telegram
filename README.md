# emersyx_telegram

Telegram receptor and resource for emersyx.

## Usage

### Plugin implementation

Source files in `emtg_impl` provide the implementation of the go plugin. They have to be built using the command:

```
go build -buildmode=plugin -o emersyx_telegram.so emtg_impl/*
```

### Plugin interface

Source files in `emtg` provide the interface to go plugin. This package has to be imported into projects which use the
plugin. The function which must be used to create new `TelegramBot` instances (which implement the `emtg.TelegramBot`
interface) is:

```
func NewTelegramBot(apiToken string) (emtg.TelegramBot, error)
```

## Credits

The underlying implementation is provided by [go-telegram-bot-api/telegram-bot-api][1].

[1]: https://github.com/go-telegram-bot-api/telegram-bot-api
