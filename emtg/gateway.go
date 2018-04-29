package main

import (
	"emersyx.net/emersyx/api"
	"emersyx.net/emersyx/api/tgapi"
	"errors"
	"github.com/BurntSushi/toml"
	"time"
)

// TelegramGateway is the type which defines a tgapi.TelegramGateway implementation, namely a Telegram resource and
// receptor for the emersyx platform.
type TelegramGateway struct {
	identifier     string
	core           api.Core
	log            *api.EmersyxLogger
	apiToken       string
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

// GetIdentifier returns the identifier of this gateway.
func (gw *TelegramGateway) GetIdentifier() string {
	return gw.identifier
}

// GetEventsOutChannel returns the api.Event channel through which emersyx events are pushed by this gateway. This
// function is required to implement the api.Receptor interface.
func (gw *TelegramGateway) GetEventsOutChannel() <-chan api.Event {
	return (<-chan api.Event)(gw.updates)
}

// NewPeripheral creates a new TelegramGateway instances based on the options given as argument.
func NewPeripheral(opts api.PeripheralOptions) (api.Peripheral, error) {
	var err error

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

	gw.identifier = opts.Identifier
	gw.core = opts.Core
	gw.log, err = api.NewEmersyxLogger(opts.LogWriter, opts.Identifier, opts.LogLevel)
	if err != nil {
		return nil, errors.New("could not create a bare logger")
	}

	// apply the extended options from the config file
	config := new(telegramGatewayConfig)
	if _, err = toml.DecodeFile(opts.ConfigPath, config); err != nil {
		return nil, err
	}
	if err = config.validate(); err != nil {
		return nil, err
	}
	config.apply(gw)

	// start polling the Telegram servers for updates
	gw.startPollingUpdates()

	return gw, nil
}
