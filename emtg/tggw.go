package main

import (
	"emersyx.net/emersyx_apis/emcomapi"
	"emersyx.net/emersyx_apis/emtgapi"
	"emersyx.net/emersyx_telegram/tgbotapi"
	"encoding/json"
	"errors"
	"time"
)

type apiResponse struct {
	OK     bool            `json:"ok"`
	Result json.RawMessage `json:"result"`
}

// TelegramGateway is the type which defines a emtgapi.TelegramGateway implementation, namely a Telegram resource and
// receptor for the emersyx platform.
type TelegramGateway struct {
	identifier     string
	updatesLimit   uint
	updatesTimeout uint
	updatesAllowed []string
	updates        chan emcomapi.Event
}

// startPollingUpdates start the process of calling the getUpdates method from the Telegram Bot API, converting the data
// to emtgapi.EUpdate objects and pushing them through the events channel of the TelegramGateway instance.
func (gw TelegramGateway) startPollingUpdates() {
	go func() {
		var offset int64
		for {
			updates, err := gw.getUpdates(offset)
			if err != nil {
				// if an error occurs, wait for 5 seconds
				time.Sleep(5)
				continue
			}
			for _, update := range updates {
				eupdate := emtgapi.EUpdate{
					Update: update,
					Source: gw.identifier,
				}
				gw.updates <- eupdate
			}
			// if we got any new updates, we need to acknowledge them
			if len(updates) > 0 {
				// the next offset value is the highest current value plus 1
				offset = updates[len(updates)-1].UpdateID + 1
			}
		}
	}()
}

// getUpdates performs calls to the getUpdates method of the Telegram Bot API and converts the data into emtgapi.Update
// instances.
func (gw TelegramGateway) getUpdates(offset int64) ([]emtgapi.Update, error) {
	var apiresp apiResponse
	var updates []emtgapi.Update

	params := (gw.NewTelegramParameters()).(*TelegramParameters)
	params.Offset(offset)
	params.Limit(gw.updatesLimit)
	params.Timeout(gw.updatesTimeout)

	resp, err := tgbotapi.GetUpdates(params.values)
	if err != nil {
		return updates, err
	}

	err = json.Unmarshal([]byte(resp), &apiresp)
	if err != nil {
		return updates, err
	}

	if apiresp.OK == false {
		return updates, errors.New("the ok field in the Bot API response is false")
	}

	err = json.Unmarshal(apiresp.Result, &updates)
	if err != nil {
		return updates, err
	}

	return updates, nil
}

// NewTelegramGateway creates new TelegramGateway instances based on the options given as argument.
func NewTelegramGateway(options ...func(emtgapi.TelegramGateway) error) (emtgapi.TelegramGateway, error) {
	gw := new(TelegramGateway)

	// create the Updates channel
	gw.updates = make(chan emcomapi.Event)

	// initialize default values for options
	gw.updatesLimit = 100
	gw.updatesTimeout = 60

	// apply the options received as argument
	for _, option := range options {
		err := option(gw)
		if err != nil {
			return nil, err
		}
	}

	// start polling the Telegram servers for updates
	gw.startPollingUpdates()

	return gw, nil
}
