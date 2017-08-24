package emtg

import (
    "errors"
    "plugin"
)

type TelegramBot interface {
    SendMessage(interface{}, string) (Message, error)
    GetEventsChannel() chan interface{}
}

func NewTelegramBot(tg_plugin *plugin.Plugin, apiToken string) (TelegramBot, error) {
    if tg_plugin == nil {
        return nil, errors.New("Invalid Telegram plugin handle.")
    }

    f, err := tg_plugin.Lookup("NewTelegramBot")
    if err != nil {
        return nil, errors.New("Telegram plugin does not have the NewTelegramBot symbol.")
    }

    var fc func(string) (TelegramBot, error)
    fc = f.(func(string) (TelegramBot, error))
    if err != nil {
        return nil, errors.New("The NewTelegramBot symbol does not have the correct signature.")
    }

    bot, err := fc(apiToken)
    return bot, err
}
