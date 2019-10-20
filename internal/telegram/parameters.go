package main

import (
	"emersyx.net/common/pkg/api/telegram"
	"errors"
	"net/url"
	"strconv"
)

// paramVals asserts the type of the telegram.Parameters argument to parameters. If the assertion is valid, then the
// function returns the values field.
func paramVals(params telegram.Parameters) (url.Values, error) {
	cparams, ok := params.(*parameters)
	if ok == false {
		return nil, errors.New("unsuppored telegram.Parameters implementation")
	}
	return cparams.values, nil
}

// parameters is the type used to set values when making calls to the Telegram Bot API.
type parameters struct {
	values url.Values
}

// Offset sets the value for the offset parameter when making a call to any Telegram Bot API method.
func (p *parameters) Offset(value int64) error {
	p.values.Set("offset", strconv.FormatInt(value, 10))
	return nil
}

// Limit sets the value for the limit parameter when making a call to any Telegram Bot API method.
func (p *parameters) Limit(value uint) error {
	p.values.Set("limit", strconv.FormatUint(uint64(value), 10))
	return nil
}

// Timeout sets the value for the timeout parameter when making a call to any Telegram Bot API method.
func (p *parameters) Timeout(value uint) error {
	p.values.Set("timeout", strconv.FormatUint(uint64(value), 10))
	return nil
}

// AllowedUpdates sets the value for the allowed_updates parameter when making a call to any Telegram Bot API method.
func (p *parameters) AllowedUpdates(values ...string) error {
	for _, value := range values {
		if len(value) > 0 {
			p.values.Add("allowed_updates", value)
		}
	}
	return nil
}

// ChatID sets the value for the chat_id parameter when making a call to any Telegram Bot API method.
func (p *parameters) ChatID(value string) error {
	if len(value) == 0 {
		return errors.New("the chat_id value cannot have zero length")
	}
	p.values.Set("chat_id", value)
	return nil
}

// Text sets the value for the text parameter when making a call to any Telegram Bot API method.
func (p *parameters) Text(value string) error {
	if len(value) == 0 {
		return errors.New("the text value cannot have zero length")
	}
	p.values.Set("text", value)
	return nil
}

// ParseMode sets the value for the parse_mode parameter when making a call to any Telegram Bot API method.
func (p *parameters) ParseMode(value string) error {
	if len(value) == 0 {
		return errors.New("the parse_mode value cannot have zero length")
	}
	p.values.Set("parse_mode", value)
	return nil
}

// DisableWebPagePreview sets the value for the disable_web_page_preview parameter when making a call to any Telegram Bot API method.
func (p *parameters) DisableWebPagePreview(value bool) error {
	p.values.Set("disable_web_page_preview", strconv.FormatBool(value))
	return nil
}

// DisableNotification sets the value for the disable_notification parameter when making a call to any Telegram Bot API method.
func (p *parameters) DisableNotification(value bool) error {
	p.values.Set("disable_notification", strconv.FormatBool(value))
	return nil
}

// ReplyToMessageID sets the value for the reply_to_message_id parameter when making a call to any Telegram Bot API method.
func (p *parameters) ReplyToMessageID(value int64) error {
	p.values.Set("reply_to_message_id", strconv.FormatInt(value, 10))
	return nil
}

// ReplyMarkup sets the value for the reply_markup parameter when making a call to any Telegram Bot API method.
func (p *parameters) ReplyMarkup(value string) error {
	if len(value) == 0 {
		return errors.New("the reply_markup value cannot have zero length")
	}
	p.values.Set("reply_markup", value)
	return nil
}
