package tgbotapi

import (
	"flag"
	"net/url"
	"os"
	"strconv"
	"testing"
)

var token = flag.String("apitoken", "", "Telegram BOT API token")
var recvID = flag.String("recvid", "", "Receiver of test messages")
var updateOffset = flag.Int64("updoffset", 0, "Value for the offset parameter when calling the getUpdates method.")

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
	params.Add("offset", strconv.FormatInt(*updateOffset, 10))
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
	// or simply obtain a token for the Telegram Bot API, send a message to a group that the bot is part of, and go to
	// https://api.telegram.org/bot<yourtoken>/getUpdates
	// and you will see JSON formatted data containing the chat_id field
	params := url.Values{}
	params.Add("chat_id", *recvID)
	params.Add("text", "hello world!")

	resp, err := SendMessage(params)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(resp)
}
