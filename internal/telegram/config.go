package main

import (
	"errors"
)

// config is the struct for holding Telegram gateway configuration values, loaded from the toml configuration file.
type config struct {
	APIToken       *string   `toml:"api_token"`
	UpdatesLimit   *uint     `toml:"updates_limit"`
	UpdatesTimeout *uint     `toml:"updates_timeout"`
	UpdatesAllowed *[]string `toml:"updates_allowed"`
}

// validate checks the values loaded from the toml configuration file. If any value is found to be invalid, then an
// error is returned.
func (c *config) validate() error {
	if c.APIToken == nil {
		return errors.New("telegram Bot API token not set")
	}
	if c.UpdatesLimit == nil {
		return errors.New("telegram Bot updates limit not set")
	}
	if c.UpdatesTimeout == nil {
		return errors.New("telegram Bot updates timeout not set")
	}
	if c.UpdatesAllowed == nil {
		return errors.New("telegram Bot list of allowed updates not set")
	}
	if len(*c.APIToken) == 0 {
		return errors.New("telegram Bot API token cannot have zero length")
	}
	if *c.UpdatesLimit < 1 || *c.UpdatesLimit > 100 {
		return errors.New("the updates limit can only be between 1 and 100")
	}
	return nil
}

// apply sets the values loaded from the toml configuration file into the gateway object received as argument.
func (c *config) apply(gw *gateway) error {
	if err := gw.setAPIToken(*c.APIToken); err == nil {
		return err
	}
	gw.updatesLimit = *c.UpdatesLimit
	gw.updatesTimeout = *c.UpdatesTimeout
	for _, t := range *c.UpdatesAllowed {
		gw.updatesAllowed = append(gw.updatesAllowed, t)
	}
	return nil
}
