package main

import(
    "flag"
    "testing"
    "fmt"
    "emersyx.net/emersyx_telegram/emtg"
)

var apiToken *string = flag.String("apitoken", "", "Telegram BOT API token")
var recvID *int64 = flag.Int64("recvid", 0, "Receiver of test messages")
var bot emtg.TelegramBot
var botInitialized bool = false

func initializeBot() {
    if botInitialized == false {
        botInitialized = true
        lbot, err := NewTelegramBot(*apiToken)
        if err != nil {
            fmt.Println(err)
        } else {
            bot = lbot
        }
    }
}

func TestSendMessage(t *testing.T) {
    initializeBot()

    /* in order to obtain the identifier of a user, channel or group, read
     * https://github.com/GabrielRF/telegram-id#web-channel-id
     */
    _, err := bot.SendMessage(*recvID, "hello world!")

    if err != nil {
        fmt.Println(err)
    }
}

func TestGetUpdates(t *testing.T) {
    k := 0
    for up := range bot.GetEventsChannel() {
        cup := up.(*emtg.Update)
        fmt.Println(cup.Message.From.Username + ": " + cup.Message.Text)
        k++
        if k == 10 {
            break
        }
    }
}
