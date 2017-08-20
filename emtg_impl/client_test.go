package main

import(
    "flag"
    "testing"
    "fmt"
)

var apiToken *string = flag.String("apitoken", "", "Telegram BOT API token")
var recvID *int64 = flag.Int64("recvid", 0, "Receiver of test messages")

func TestApiKey(t *testing.T) {
    _, err := NewTelegramBot(*apiToken)

    if err != nil {
        fmt.Println(err)
    }
}

func TestSendMessage(t *testing.T) {
    bot, err := NewTelegramBot(*apiToken)
    if err != nil {
        fmt.Println(err)
        return
    }

    /* in order to obtain the identifier of a user, channel or group, read
     * https://github.com/GabrielRF/telegram-id#web-channel-id
     */
    _, err = bot.SendMessage(*recvID, "hello world!")

    if err != nil {
        fmt.Println(err)
    }
}
