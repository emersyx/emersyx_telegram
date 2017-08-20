package main

import(
    "emersyx.net/emersyx_telegram/emtg"
    "errors"
    "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramBot struct {
    api *tgbotapi.BotAPI
    Updates chan interface{}
}

func NewTelegramBot(apiToken string) (emtg.TelegramBot, error) {
    var err error

    bot := new(TelegramBot)
    bot.Updates = make(chan interface{})
    bot.api, err = tgbotapi.NewBotAPI(apiToken)
    if err != nil {
        return nil, err
    }

    ucfg := tgbotapi.NewUpdate(0)
    ucfg.Timeout = 60
    uch, err := bot.api.GetUpdatesChan(ucfg)
    if err != nil {
        return nil, err
    }

    go func() {
        var up tgbotapi.Update
        for true {
            up = <- uch
            bot.Updates <- up
        }
    }()

    return bot, nil
}

func (bot *TelegramBot) GetEventsChannel() chan interface{} {
    return bot.Updates
}

func (bot *TelegramBot) SendMessage(chatID interface{}, text string) (emtg.Message, error) {
    var smsg tgbotapi.MessageConfig

    switch v := chatID.(type) {
    case int64:
        smsg = tgbotapi.NewMessage(v, text)
    case string:
        smsg = tgbotapi.NewMessageToChannel(v, text)
    default:
        return emtg.Message{}, errors.New("invalid type for chatID argument")
    }

    rmsg, rerr := bot.api.Send(smsg)
    cmsg, cerr := convertMessage(rmsg)

    if rerr != nil {
        if cerr != nil {
            return emtg.Message{}, cerr
        } else {
            return cmsg, rerr
        }
    } else {
        if cerr != nil {
            return emtg.Message{}, cerr
        } else {
            return cmsg, nil
        }
    }
}

