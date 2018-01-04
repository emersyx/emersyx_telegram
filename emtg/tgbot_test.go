package main

import(
    "flag"
    "fmt"
    "testing"
    "os"
    "emersyx.net/emersyx_apis/emtgapi"
)

var token = flag.String("apitoken", "", "Telegram BOT API token")
var recvID = flag.String("recvid", "", "Receiver of test messages")
var updateOffset = flag.Int64("updoffset", 0, "Value for the offset parameter when calling the getUpdates method.")
var bot emtgapi.TelegramBot

func TestMain(m *testing.M) {
    var err error

    // get the command line flags
    flag.Parse()

    // generate a TelegramOptions object and set the options for the TelegramBot
    opt := NewTelegramOptions()

    // create the telegram bot
    // in this implementation, the NewTelegramBot function also makes a call to getMe
    bot, err = NewTelegramBot(
        opt.APIToken(*token),
        opt.Identifier("emersyx-tgbot-test"),
    )

    if err != nil {
        fmt.Println( err.Error() )
    } else {
        // run the tests
        os.Exit(m.Run())
    }
}

func TestGetMe(t *testing.T) {
    u, err := bot.GetMe()
    if err != nil {
        t.Log( err.Error() )
        t.Fail()
        return
    }

    t.Log("TestGetMe")
    t.Log("ID           ", u.ID)
    t.Log("IsBot        ", u.IsBot)
    t.Log("FirstName    ", u.FirstName)
    t.Log("LastName     ", u.LastName)
    t.Log("Username     ", u.Username)
    t.Log("LanguageCode ", u.LanguageCode)
    t.Log("-----")
}

func TestSendMessage(t *testing.T) {
    params := NewTelegramParameters()
    params.ChatID(*recvID)
    params.Text("hello world! hello from *emersyx*!")
    params.ParseMode("Markdown")
    m, err := bot.SendMessage(params)
    if err != nil {
        t.Log( err.Error() )
        t.Fail()
        return
    }

    t.Log("TestSendMessage")
    t.Log("MessageID    ", m.MessageID)
    t.Log("From         ", m.From)
    t.Log("Date         ", m.Date)
    t.Log("Chat         ", m.Chat)
    t.Log("-----")
}

func TestGetUpdates(t *testing.T) {
    eventsChannel := bot.GetEventsChannel()
    e := <- eventsChannel

    if e.GetSourceIdentifier() != "emersyx-tgbot-test" {
        t.Fail()
        return
    }

    eu, ok := e.(emtgapi.EUpdate)
    if ok == false {
        t.Log("The event type is not emtgapi.EUpdate.")
        t.Fail()
        return
    }

    t.Log("MessageID    ", eu.Update.Message.MessageID)
    t.Log("From         ", eu.Update.Message.From)
    t.Log("Date         ", eu.Update.Message.Date)
    t.Log("Chat         ", eu.Update.Message.Chat)
    t.Log("-----")
}
