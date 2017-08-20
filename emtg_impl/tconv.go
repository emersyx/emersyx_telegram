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
