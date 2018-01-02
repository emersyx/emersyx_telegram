package tgbotapi

import(
    "flag"
    "net/url"
    "os"
    "testing"
)

var token *string = flag.String("apitoken", "", "Telegram BOT API token")
var recvID *string = flag.String("recvid", "", "Receiver of test messages")

func TestMain(m *testing.M) {
    // get the command line flags
    flag.Parse()

    // initialize the low level Telegram Bot API library
    Initialize(*token)

    // run the tests
    os.Exit(m.Run())
}

func TestGetUpdates(t *testing.T) {
    params := url.Values{}
    //params.Add("offset", "80854227")
    resp, err := GetUpdates(params)
    if err != nil {
        t.Fatal(err.Error())
    }
    t.Log(resp)
}

func TestGetMe(t *testing.T) {
    resp, err := GetMe()
    if err != nil {
        t.Fatal(err.Error())
    }
    t.Log(resp)
}

func TestSendMessage(t *testing.T) {
    // in order to obtain the identifier of a user, channel or group, read
    // https://github.com/GabrielRF/telegram-id#web-channel-id
    params := url.Values{}
    params.Add("chat_id", *recvID)
    params.Add("text", "hello world!")

    resp, err := SendMessage(params)
    if err != nil {
        t.Fatal(err.Error())
    }
    t.Log(resp)
}
