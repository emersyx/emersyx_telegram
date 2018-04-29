package main

import (
	"errors"
)

// telegramGatewayConfig is the struct for holding Telegram gateway configuration values, loaded from the configuration
// file.
type telegramGatewayConfig struct {
	APIToken       *string   `toml:"api_token"`
	UpdatesLimit   *uint     `toml:"updates_limit"`
	UpdatesTimeout *uint     `toml:"updates_timeout"`
	UpdatesAllowed *[]string `toml:"updates_allowed"`
}

// validate checks the values loaded from the toml configuration file. If any value is found to be invalid, then an
// error is returned.
func (cfg *telegramGatewayConfig) validate() error {
	if cfg.APIToken == nil {
		return errors.New("telegram Bot API token not set")
	}
	if cfg.UpdatesLimit == nil {
		return errors.New("telegram Bot updates limit not set")
	}
	if cfg.UpdatesTimeout == nil {
		return errors.New("telegram Bot updates timeout not set")
	}
	if cfg.UpdatesAllowed == nil {
		return errors.New("telegram Bot list of allowed updates not set")
	}
	if len(*cfg.APIToken) == 0 {
		return errors.New("telegram Bot API token cannot have zero length")
	}
	if *cfg.UpdatesLimit < 1 || *cfg.UpdatesLimit > 100 {
		return errors.New("the updates limit can only be between 1 and 100")
	}
	return nil
}

// apply sets the values loaded from the toml configuration file into the ircGateway object received as argument.
func (cfg *telegramGatewayConfig) apply(gw *TelegramGateway) error {
	if err := gw.setAPIToken(*cfg.APIToken); err == nil {
		return err
	}
	gw.updatesLimit = *cfg.UpdatesLimit
	gw.updatesTimeout = *cfg.UpdatesTimeout
	for _, t := range *cfg.UpdatesAllowed {
		gw.updatesAllowed = append(gw.updatesAllowed, t)
	}
	return nil
}
