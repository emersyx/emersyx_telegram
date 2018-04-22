package main

import (
	"emersyx.net/emersyx/api"
	"emersyx.net/emersyx/api/tgapi"
	"emersyx.net/emersyx_telegram/tgbotapi"
	"encoding/json"
	"errors"
)

// GetIdentifier returns the identifier of this receptor.
func (gw *TelegramGateway) GetIdentifier() string {
	return gw.identifier
}

// GetMe performs a call to the getMe method of the Telegram Bot API.
func (gw *TelegramGateway) GetMe() (tgapi.User, error) {
	apiresp := apiResponse{}
	user := tgapi.User{}

	resp, err := tgbotapi.GetMe()
	if err != nil {
		return user, err
	}

	err = json.Unmarshal([]byte(resp), &apiresp)
	if err != nil {
		return user, err
	}

	if apiresp.OK == false {
		return user, errors.New("the ok field in the Bot API response is false")
	}

	err = json.Unmarshal(apiresp.Result, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// SendMessage performs a call to the sendMessage method of the Telegram Bot API.
func (gw *TelegramGateway) SendMessage(params tgapi.TelegramParameters) (tgapi.Message, error) {
	apiresp := apiResponse{}
	msg := tgapi.Message{}

	cparams, ok := params.(*TelegramParameters)
	if ok == false {
		return msg, errors.New("unsuppored TelegramParameters implementation")
	}

	resp, err := tgbotapi.SendMessage(cparams.values)
	if err != nil {
		return msg, err
	}

	err = json.Unmarshal([]byte(resp), &apiresp)
	if err != nil {
		return msg, err
	}

	if apiresp.OK == false {
		return msg, errors.New("the ok field in the Bot API response is false")
	}

	err = json.Unmarshal(apiresp.Result, &msg)
	if err != nil {
		return msg, err
	}

	return msg, nil
}

// NewTelegramParameters creates and returns a new TelegramParameters object. This object can then be used to configure
// parameters when performing calls to the Telegram Bot API (e.g. see the TelegramGateway.SendMessage method).
func (gw *TelegramGateway) NewTelegramParameters() tgapi.TelegramParameters {
	params := new(TelegramParameters)
	params.values = make(map[string][]string)
	return params
}

// GetEventsOutChannel returns the api.Event channel through which emersyx events are pushed by this gateway.
func (gw *TelegramGateway) GetEventsOutChannel() <-chan api.Event {
	return (<-chan api.Event)(gw.updates)
}

// GetEventsInChannel returns the api.CoreEvent channel through which core events are received by the gateway
// instance.
func (gw *TelegramGateway) GetEventsInChannel() chan<- api.Event {
	return nil
}
