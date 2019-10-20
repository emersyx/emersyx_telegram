package main

import (
	"emersyx.net/common/pkg/api"
	"emersyx.net/common/pkg/api/telegram"
	"emersyx.net/common/pkg/log"
	"flag"
	"fmt"
	"os"
	"testing"
)

var recvID = flag.String("recvid", "", "Receiver of test messages")
var updateOffset = flag.Int64("updoffset", 0, "Value for the offset parameter when calling the getUpdates method.")
var conffile = flag.String("conffile", "", "path to toml configuration file")
var gw telegram.Gateway

func TestMain(m *testing.M) {
	var err error
	var ok bool

	// get the command line flags
	flag.Parse()

	// create a new telegramGateway
	peripheral, err := NewPeripheral(
		api.PeripheralOptions{
			Identifier: "telegram-gateway-test",
			ConfigPath: *conffile,
			LogWriter:  os.Stdout,
			LogLevel:   log.ELDebug,
		},
	)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	gw, ok = peripheral.(*gateway)
	if ok != true {
		fmt.Println("invalid peripheral type")
	} else {
		// run the tests
		os.Exit(m.Run())
	}
}

func TestGetMe(t *testing.T) {
	u, err := gw.GetMe()
	if err != nil {
		t.Log(err.Error())
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
	params := gw.NewParameters()
	params.ChatID(*recvID)
	params.Text("hello world! hello from *emersyx*!")
	params.ParseMode("Markdown")
	m, err := gw.SendMessage(params)
	if err != nil {
		t.Log(err.Error())
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
	eventsChannel := gw.GetEventsOutChannel()
	e := <-eventsChannel

	if e.GetSourceIdentifier() != "telegram-gateway-test" {
		t.Fail()
		return
	}

	u, ok := e.(telegram.Update)
	if ok == false {
		t.Log("The event type is not telegram.Update.")
		t.Fail()
		return
	}

	t.Log("MessageID    ", u.Message.MessageID)
	t.Log("From         ", u.Message.From)
	t.Log("Date         ", u.Message.Date)
	t.Log("Chat         ", u.Message.Chat)
	t.Log("-----")
}
