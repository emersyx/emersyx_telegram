package main

import (
	"emersyx.net/emersyx/api"
	"errors"
	"github.com/BurntSushi/toml"
	"io"
)

// telegramOptions implements the api.PeripheralOptions interface. Each method returns a function, which applies a
// specific configuration to a TelegramGateway object.
// TODO make this private
type telegramOptions struct {
}

// Logging sets the io.Writer instance to write logging messages to and the verbosity level.
func (o telegramOptions) Logging(writer io.Writer, level uint) func(api.Peripheral) error {
	return func(peripheral api.Peripheral) error {
		if writer == nil {
			return errors.New("writer argument cannot be nil")
		}
		cgw := assertTelegramGateway(peripheral)
		cgw.log.SetOutput(writer)
		cgw.log.SetLevel(level)
		return nil
	}
}

// Identifier sets the receptor identifier value for the Telegram gateway.
func (o telegramOptions) Identifier(id string) func(api.Peripheral) error {
	return func(peripheral api.Peripheral) error {
		if len(id) == 0 {
			return errors.New("identifier value cannot have zero length")
		}
		gw := assertTelegramGateway(peripheral)
		gw.identifier = id
		gw.log.SetComponentID(id)
		return nil
	}
}

// ConfigPath loads the toml configuration file and validates the contents. If valid, the contents are applied to the
// Telegram gateway.
func (o telegramOptions) ConfigPath(path string) func(api.Peripheral) error {
	return func(peripheral api.Peripheral) error {
		config := new(telegramGatewayConfig)
		_, err := toml.DecodeFile(path, config)
		if err != nil {
			return err
		}
		if err := config.validate(); err != nil {
			return err
		}
		gw := assertTelegramGateway(peripheral)
		// TODO check return error of config.apply
		config.apply(gw)
		return nil
	}
}

// Core sets the api.Core instance which provides services to the Telegram gateway.
func (o telegramOptions) Core(core api.Core) func(api.Peripheral) error {
	return func(peripheral api.Peripheral) error {
		if core == nil {
			return errors.New("core argument cannot be nil")
		}
		cgw := assertTelegramGateway(peripheral)
		cgw.core = core
		return nil
	}
}

// assertTelegramGateway tries to make a type assertion on the peripheral argument, to the *TelegramGateway type. If the
// type assertion fails, then panic() is called. A call to recover() is in the applyOptions function.
func assertTelegramGateway(peripheral api.Peripheral) *TelegramGateway {
	gw, ok := peripheral.(*TelegramGateway)
	if ok == false {
		panic("unsupported TelegramGateway implementation")
	}
	return gw
}

// applyOptions executes the functions provided as the options argument with Telegram gateway as argument. The
// implementation calls recover() in order to stop panicking, which may be caused by the call to panic() within the
// assertTelegramGateway function. assertTelegramGateway is used by functions returned by ircOptions.
func applyOptions(peripheral api.Peripheral, options ...func(api.Peripheral) error) (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = r.(error)
		}
	}()

	for _, option := range options {
		if e = option(peripheral); e != nil {
			return
		}
	}

	return
}

// NewPeripheralOptions createa a new telegramOptions instance and returns a pointer to it.
func NewPeripheralOptions() api.PeripheralOptions {
	return new(telegramOptions)
}
