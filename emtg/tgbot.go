package main

import(
    "encoding/json"
    "errors"
    "time"
    "emersyx.net/emersyx_apis/emcomapi"
    "emersyx.net/emersyx_apis/emtgapi"
    "emersyx.net/emersyx_telegram/tgbotapi"
)

type apiResponse struct {
    OK      bool            `json:"ok"`
    Result  json.RawMessage `json:"result"`
}

type TelegramBot struct {
    identifier string
    updatesLimit uint
    updatesTimeout uint
    updatesAllowed []string
    updates chan emcomapi.Event
}

func (bot TelegramBot) startPollingUpdates() {
    go func() {
        var offset int64
        for {
            updates, err := bot.getUpdates(offset)
            if err != nil {
                // if an error occurs, wait for 5 seconds
                time.Sleep(5)
                continue
            }
            for _, update := range updates {
                eupdate := emtgapi.EUpdate {
                    Update : update,
                    Source : bot.identifier,
                }
                bot.updates <- eupdate
            }
            // the next offset value is the highest current value plus 1
            offset = updates[ len(updates) - 1 ].UpdateID + 1
        }
    }()
}

func (bot TelegramBot) getUpdates(offset int64) ([]emtgapi.Update, error) {
    var apiresp apiResponse
    var updates []emtgapi.Update

    params := (NewTelegramParameters()).(*TelegramParameters)
    params.Offset(offset)
    params.Limit(bot.updatesLimit)
    params.Timeout(bot.updatesTimeout)

    resp, err := tgbotapi.GetUpdates(params.values)
    if err != nil {
        return updates, err
    }

    err = json.Unmarshal( []byte(resp), &apiresp )
    if err != nil {
        return updates, err
    }

    if apiresp.OK == false {
        return updates, errors.New("The ok field in the Bot API response is false.")
    }

    err = json.Unmarshal(apiresp.Result, &updates)
    if err != nil {
        return updates, err
    }

    return updates, nil
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
    bot.startPollingUpdates()

    return bot, nil
}
