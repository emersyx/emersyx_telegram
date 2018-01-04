package main

import(
    "encoding/json"
    "errors"
    "emersyx.net/emersyx_apis/emcomapi"
    "emersyx.net/emersyx_apis/emtgapi"
    "emersyx.net/emersyx_telegram/tgbotapi"
)

// GetIdentifier returns the identifier of this receptor.
func (bot *TelegramBot) GetIdentifier() string {
    return bot.identifier
}

// GetEventsChannel returns get emcomapi.Event channel through which this receptor pushes emersyx events.
func (bot *TelegramBot) GetEventsChannel() chan emcomapi.Event {
    return bot.updates
}

// GetMe performs a call to the getMe method of the Telegram Bot API.
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
        return user, errors.New("the ok field in the Bot API response is false")
    }

    err = json.Unmarshal(apiresp.Result, &user)
    if err != nil {
        return user, err
    }

    return user, nil
}

// SendMessage performs a call to the sendMessage method of the Telegram Bot API.
func (bot *TelegramBot) SendMessage(params emtgapi.TelegramParameters) (emtgapi.Message, error) {
    apiresp := apiResponse{}
    msg := emtgapi.Message{}

    cparams, ok := params.(*TelegramParameters)
    if ok == false {
        return msg, errors.New("unsuppored TelegramParameters implementation")
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
        return msg, errors.New("the ok field in the Bot API response is false")
    }

    err = json.Unmarshal(apiresp.Result, &msg)
    if err != nil {
        return msg, err
    }

    return msg, nil
}
