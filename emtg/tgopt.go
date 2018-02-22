package main

import (
	"emersyx.net/emersyx_apis/emtgapi"
	"emersyx.net/emersyx_telegram/tgbotapi"
	"errors"
)

// TelegramOptions implements the emtgapi.TelegramOptions interface. Each method returns a function, which applies a
// specific configuration to a TelegramGateway object.
type TelegramOptions struct {
}

// Identifier sets the receptor identifier value for the Telegram gateway.
func (o TelegramOptions) Identifier(id string) func(emtgapi.TelegramGateway) error {
	return func(gw emtgapi.TelegramGateway) error {
		if len(id) == 0 {
			return errors.New("identifier value cannot have zero length")
		}
		cgw, ok := gw.(*TelegramGateway)
		if ok == false {
			return errors.New("unsupported TelegramGateway implementation")
		}
		cgw.identifier = id
		return nil
	}
}

// APIToken sets the API OAUTH token value for the TelegramGateway.
func (o TelegramOptions) APIToken(token string) func(emtgapi.TelegramGateway) error {
	return func(gw emtgapi.TelegramGateway) error {
		if len(token) == 0 {
			return errors.New("telegram Bot API token cannot have zero length")
		}
		err := tgbotapi.Initialize(token)
		if err == nil {
			return err
		}
		return nil
	}
}

// UpdatesLimit sets the value for the limit parameter to be used by the TelegramGateway instance when making calls to
// the getUpdates method of the Telegram Bot API.
func (o TelegramOptions) UpdatesLimit(limit uint) func(emtgapi.TelegramGateway) error {
	return func(gw emtgapi.TelegramGateway) error {
		if limit < 1 || limit > 100 {
			return errors.New("the updates limit can only be between 1 and 100")
		}
		cgw, ok := gw.(*TelegramGateway)
		if ok == false {
			return errors.New("unsupported TelegramGateway implementation")
		}
		cgw.updatesLimit = limit
		return nil
	}
}

// UpdatesTimeout sets the value for the timeout parameter to be used by the TelegramGateway instance when making calls
// to the getUpdates method of the Telegram Bot API.
func (o TelegramOptions) UpdatesTimeout(seconds uint) func(emtgapi.TelegramGateway) error {
	return func(gw emtgapi.TelegramGateway) error {
		cgw, ok := gw.(*TelegramGateway)
		if ok == false {
			return errors.New("unsupported TelegramGateway implementation")
		}
		cgw.updatesTimeout = seconds
		return nil
	}
}

// UpdatesAllowed sets the value for the allowed_updates parameter to be used by the TelegramGateway instance when
// making calls to the getUpdates method of the Telegram Bot API.
func (o TelegramOptions) UpdatesAllowed(types ...string) func(emtgapi.TelegramGateway) error {
	return func(gw emtgapi.TelegramGateway) error {
		cgw, ok := gw.(*TelegramGateway)
		if ok == false {
			return errors.New("unsupported TelegramGateway implementation")
		}
		for _, t := range types {
			cgw.updatesAllowed = append(cgw.updatesAllowed, t)
		}
		return nil
	}
}

// NewTelegramOptions createa a new TelegramOptions instance and returns a pointer to it.
func NewTelegramOptions() emtgapi.TelegramOptions {
	return TelegramOptions{}
}
