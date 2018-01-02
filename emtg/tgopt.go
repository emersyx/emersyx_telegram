package main

import(
    "errors"
    "emersyx.net/emersyx_apis/emtgapi"
    "emersyx.net/emersyx_telegram/tgbotapi"
)

type TelegramOptions struct {
}

func (o TelegramOptions) APIToken(token string) func(emtgapi.TelegramBot) error {
    return func(bot emtgapi.TelegramBot) error {
        if len(token) == 0 {
            return errors.New("Telegram Bot API token cannot have zero length.")
        }
        err := tgbotapi.Initialize(token)
        if err == nil {
            return err
        }
        return nil
    }
}

func (o TelegramOptions) UpdatesLimit(limit uint) func(emtgapi.TelegramBot) error {
    return func(bot emtgapi.TelegramBot) error {
        if limit < 1 || limit > 100 {
            return errors.New("The updates limit can only be between 1 and 100.")
        }
        cbot, ok := bot.(*TelegramBot)
        if ok == false {
            return errors.New("Unsupported TelegramBot implementation.")
        }
        cbot.updatesLimit = limit
        return nil
    }
}

func (o TelegramOptions) UpdatesTimeout(seconds uint) func(emtgapi.TelegramBot) error {
    return func(bot emtgapi.TelegramBot) error {
        cbot, ok := bot.(*TelegramBot)
        if ok == false {
            return errors.New("Unsupported TelegramBot implementation.")
        }
        cbot.updatesTimeout = seconds
        return nil
    }
}

func (o TelegramOptions) UpdatesAllowed(types ...string) func(emtgapi.TelegramBot) error {
    return func(bot emtgapi.TelegramBot) error {
        cbot, ok := bot.(*TelegramBot)
        if ok == false {
            return errors.New("Unsupported TelegramBot implementation.")
        }
        for _, t := range types {
            cbot.updatesAllowed = append(cbot.updatesAllowed, t)
        }
        return nil
    }
}
