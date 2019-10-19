package main

import (
	"emersyx.net/emersyx/pkg/api"
	"github.com/BurntSushi/toml"
	"time"
)

// gateway is the type which defines a telegram.Gateway implementation, namely a Telegram resource and receptor for the
// emersyx platform.
type gateway struct {
	api.PeripheralBase
	apiToken       string
	updatesLimit   uint
	updatesTimeout uint
	updatesAllowed []string
	updates        chan api.Event
}

// startPollingUpdates start the process of calling the getUpdates method from the Telegram Bot API pushing
// telegram.Update objects through the events channel of the gateway instance.
func (gw *gateway) startPollingUpdates() {
	go func() {
		var offset int64
		for {
			updates, err := gw.getUpdates(offset)
			if err != nil {
				gw.Log.Errorf("received error while getting updates %s\n", err.Error())
				// if an error occurs, wait for 5 seconds
				time.Sleep(5)
				continue
			}
			gw.Log.Debugf("received %d update(s)\n", len(updates))
			for _, update := range updates {
				update.Source = gw.Identifier
				gw.updates <- update
			}
			// if we got any new updates, we need to acknowledge them
			if len(updates) > 0 {
				gw.Log.Debugln("acknowledging updates to the Telegram Bot back-end")
				// the next offset value is the highest current value plus 1
				offset = updates[len(updates)-1].UpdateID + 1
			}
		}
	}()
}

// GetIdentifier returns the identifier of this gateway.
func (gw *gateway) GetIdentifier() string {
	return gw.Identifier
}

// GetEventsOutChannel returns the api.Event channel through which emersyx events are pushed by this gateway. This
// function is required to implement the api.Receptor interface.
func (gw *gateway) GetEventsOutChannel() <-chan api.Event {
	return (<-chan api.Event)(gw.updates)
}

// NewPeripheral creates a new gateway instances based on the options given as argument.
func NewPeripheral(opts api.PeripheralOptions) (api.Peripheral, error) {
	var err error

	// create a new gateway and initialize the base
	gw := new(gateway)
	gw.InitializeBase(opts)

	// create the Updates channel
	gw.updates = make(chan api.Event)

	// initialize default values for options
	gw.updatesLimit = 100
	gw.updatesTimeout = 60

	// apply the extended options from the config file
	c := new(config)
	if _, err = toml.DecodeFile(opts.ConfigPath, c); err != nil {
		return nil, err
	}
	if err = c.validate(); err != nil {
		return nil, err
	}
	c.apply(gw)

	// start polling the Telegram servers for updates
	gw.startPollingUpdates()

	return gw, nil
}
