package main

import(
    "encoding/json"
    "errors"
    "emersyx.net/emersyx_apis/emcomapi"
    "emersyx.net/emersyx_apis/emtgapi"
    "emersyx.net/emersyx_telegram/tgbotapi"
)

func (bot *TelegramBot) GetIdentifier() string {
    return bot.identifier
}

func (bot *TelegramBot) GetEventsChannel() chan emcomapi.Event {
    return bot.updates
}

func (bot *TelegramBot) GetMe() (emtgapi.User, error) {
    u := emtgapi.User{}

    resp, err := tgbotapi.GetMe()
    if err != nil {
        return u, err
    }

    err = json.Unmarshal( []byte(resp), &u )
    if err != nil {
        return u, err
    }

    return u, nil
}

func (bot *TelegramBot) SendMessage(params emtgapi.TelegramParameters) (emtgapi.Message, error) {
    m := emtgapi.Message{}

    cparams, ok := params.(*TelegramParameters)
    if ok == false {
        return m, errors.New("Unsuppored TelegramParameters implementation.")
    }

    resp, err := tgbotapi.SendMessage(cparams.values)
    if err != nil {
        return m, err
    }

    err = json.Unmarshal( []byte(resp), &m )
    if err != nil {
        return m, err
    }

    return m, nil
}
