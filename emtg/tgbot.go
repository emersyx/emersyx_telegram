package main

import(
    "emersyx.net/emersyx_apis/emcomapi"
    "emersyx.net/emersyx_apis/emtgapi"
)

type TelegramBot struct {
    identifier string
    updatesLimit uint
    updatesTimeout uint
    updatesAllowed []string
    updates chan emcomapi.Event
}

func (bot TelegramBot) startPolling() {
}

func NewTelegramBot(options ...func(emtgapi.TelegramBot) error) (emtgapi.TelegramBot, error) {
    bot := new(TelegramBot)

    // create the Updates channel
    bot.updates = make(chan emcomapi.Event)

    // initialize default values for options
    bot.updatesLimit = 100
    bot.updatesTimeout = 60

    // apply the options received as argument
    for _, option := range options {
        err := option(bot)
        if err != nil {
            return nil, err
        }
    }

    // start polling the Telegram servers for updates
    bot.startPolling()

    return bot, nil
}
