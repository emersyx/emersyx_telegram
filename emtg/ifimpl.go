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
    apiresp := apiResponse{}
    user := emtgapi.User{}

    resp, err := tgbotapi.GetMe()
    if err != nil {
        return user, err
    }

    err = json.Unmarshal( []byte(resp), &apiresp )
    if err != nil {
        return user, err
    }

    if apiresp.OK == false {
        return user, errors.New("The ok field in the Bot API response is false.")
    }

    err = json.Unmarshal(apiresp.Result, &user)
    if err != nil {
        return user, err
    }

    return user, nil
}

func (bot *TelegramBot) SendMessage(params emtgapi.TelegramParameters) (emtgapi.Message, error) {
    apiresp := apiResponse{}
    msg := emtgapi.Message{}

    cparams, ok := params.(*TelegramParameters)
    if ok == false {
        return msg, errors.New("Unsuppored TelegramParameters implementation.")
    }

    resp, err := tgbotapi.SendMessage(cparams.values)
    if err != nil {
        return msg, err
    }

    err = json.Unmarshal( []byte(resp), &apiresp )
    if err != nil {
        return msg, err
    }

    if apiresp.OK == false {
        return msg, errors.New("The ok field in the Bot API response is false.")
    }

    err = json.Unmarshal(apiresp.Result, &msg)
    if err != nil {
        return msg, err
    }

    return msg, nil
}
