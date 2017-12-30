package main

import(
    "emersyx.net/emersyx_telegram/emtg"
    "encoding/json"
    "github.com/go-telegram-bot-api/telegram-bot-api"
)

func convertMessage(am tgbotapi.Message) (emtg.Message, error) {
    bytes, err := json.Marshal(am)
    if err != nil {
        return emtg.Message{}, err
    }

    rm := emtg.Message{}
    err = json.Unmarshal(bytes, &rm)
    if err != nil {
        return emtg.Message{}, err
    }

    return rm, nil
}

func convertUpdate(au tgbotapi.Update) (emtg.Update, error) {
    bytes, err := json.Marshal(au)
    if err != nil {
        return emtg.Update{}, err
    }

    ru := emtg.Update{}
    err = json.Unmarshal(bytes, &ru)
    if err != nil {
        return emtg.Update{}, err
    }

    return ru, nil
}
