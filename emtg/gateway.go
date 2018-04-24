package main

import (
	"emersyx.net/emersyx/api"
	"emersyx.net/emersyx/api/tgapi"
	"emersyx.net/emersyx/log"
	"emersyx.net/emersyx_telegram/tgbotapi"
	"encoding/json"
	"errors"
	"github.com/BurntSushi/toml"
	"time"
)

type apiResponse struct {
	OK     bool            `json:"ok"`
	Result json.RawMessage `json:"result"`
}

// TelegramGateway is the type which defines a tgapi.TelegramGateway implementation, namely a Telegram resource and
// receptor for the emersyx platform.
type TelegramGateway struct {
	core           api.Core
	log            *log.EmersyxLogger
	isInitialized  bool
	identifier     string
	updatesLimit   uint
	updatesTimeout uint
	updatesAllowed []string
	updates        chan api.Event
}

// startPollingUpdates start the process of calling the getUpdates method from the Telegram Bot API, converting the data
// to tgapi.EUpdate objects and pushing them through the events channel of the TelegramGateway instance.
func (gw TelegramGateway) startPollingUpdates() {
	go func() {
		var offset int64
		for {
			updates, err := gw.getUpdates(offset)
			if err != nil {
				gw.log.Errorf("received error while getting updates %s\n", err.Error())
				// if an error occurs, wait for 5 seconds
				time.Sleep(5)
				continue
			}
			gw.log.Debugf("received %d update(s)\n", len(updates))
			for _, update := range updates {
				eupdate := tgapi.EUpdate{
					Update: update,
					Source: gw.identifier,
				}
				gw.updates <- eupdate
			}
			// if we got any new updates, we need to acknowledge them
			if len(updates) > 0 {
				gw.log.Debugln("acknowledging updates to the Telegram Bot back-end")
				// the next offset value is the highest current value plus 1
				offset = updates[len(updates)-1].UpdateID + 1
			}
		}
	}()
}

// getUpdates performs calls to the getUpdates method of the Telegram Bot API and converts the data into tgapi.Update
// instances.
func (gw TelegramGateway) getUpdates(offset int64) ([]tgapi.Update, error) {
	var apiresp apiResponse
	var updates []tgapi.Update

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

// NewPeripheral creates a new TelegramGateway instances based on the options given as argument.
func NewPeripheral(opts api.PeripheralOptions) (api.Peripheral, error) {
	// validate identifier in options
	if len(opts.Identifier) == 0 {
		return nil, errors.New("identifier cannot have 0 length")
	}

	gw := new(TelegramGateway)

	// create the Updates channel
	gw.updates = make(chan api.Event)

	// initialize default values for options
	gw.updatesLimit = 100
	gw.updatesTimeout = 60

	// generate a bare logger, to be updated via options
	if log, err := log.NewEmersyxLogger(nil, "", log.ELNone); err != nil {
		gw.log = log
		return nil, errors.New("could not create a bare logger")
	}

	// apply the options received as argument
	gw.identifier = opts.Identifier
	gw.core = opts.Core
	gw.log.SetOutput(opts.LogWriter)
	gw.log.SetLevel(opts.LogLevel)
	gw.log.SetComponentID(gw.identifier)

	// apply the extended options from the config file
	config := new(telegramGatewayConfig)
	if _, err := toml.DecodeFile(opts.ConfigPath, config); err != nil {
		return nil, err
	}
	if err := config.validate(); err != nil {
		return nil, err
	}
	config.apply(gw)

	// start polling the Telegram servers for updates
	gw.startPollingUpdates()

	return gw, nil
}
