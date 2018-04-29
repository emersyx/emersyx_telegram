package main

import (
	"emersyx.net/emersyx/api/tgapi"
	"errors"
	"net/url"
	"strconv"
)

// paramVals asserts the type of the params argument to telegramParameters. If the assertion is valid, then the function
// returns the values field.
func paramVals(params tgapi.TelegramParameters) (url.Values, error) {
	cparams, ok := params.(*telegramParameters)
	if ok == false {
		return nil, errors.New("unsuppored TelegramParameters implementation")
	}
	return cparams.values, nil
}

// telegramParameters is the type which is used to set values for parameters when making calls to the Telegram Bot API.
type telegramParameters struct {
	values url.Values
}

// Offset sets the value for the offset parameter when making a call to any Telegram Bot API method.
func (params *telegramParameters) Offset(value int64) error {
	params.values.Set("offset", strconv.FormatInt(value, 10))
	return nil
}

// Limit sets the value for the limit parameter when making a call to any Telegram Bot API method.
func (params *telegramParameters) Limit(value uint) error {
	params.values.Set("limit", strconv.FormatUint(uint64(value), 10))
	return nil
}

// Timeout sets the value for the timeout parameter when making a call to any Telegram Bot API method.
func (params *telegramParameters) Timeout(value uint) error {
	params.values.Set("timeout", strconv.FormatUint(uint64(value), 10))
	return nil
}

// AllowedUpdates sets the value for the allowed_updates parameter when making a call to any Telegram Bot API method.
func (params *telegramParameters) AllowedUpdates(values ...string) error {
	for _, value := range values {
		if len(value) > 0 {
			params.values.Add("allowed_updates", value)
		}
	}
	return nil
}

// ChatID sets the value for the chat_id parameter when making a call to any Telegram Bot API method.
func (params *telegramParameters) ChatID(value string) error {
	if len(value) == 0 {
		return errors.New("the chat_id value cannot have zero length")
	}
	params.values.Set("chat_id", value)
	return nil
}

// Text sets the value for the text parameter when making a call to any Telegram Bot API method.
func (params *telegramParameters) Text(value string) error {
	if len(value) == 0 {
		return errors.New("the text value cannot have zero length")
	}
	params.values.Set("text", value)
	return nil
}

// ParseMode sets the value for the parse_mode parameter when making a call to any Telegram Bot API method.
func (params *telegramParameters) ParseMode(value string) error {
	if len(value) == 0 {
		return errors.New("the parse_mode value cannot have zero length")
	}
	params.values.Set("parse_mode", value)
	return nil
}

// DisableWebPagePreview sets the value for the disable_web_page_preview parameter when making a call to any Telegram Bot API method.
func (params *telegramParameters) DisableWebPagePreview(value bool) error {
	params.values.Set("disable_web_page_preview", strconv.FormatBool(value))
	return nil
}

// DisableNotification sets the value for the disable_notification parameter when making a call to any Telegram Bot API method.
func (params *telegramParameters) DisableNotification(value bool) error {
	params.values.Set("disable_notification", strconv.FormatBool(value))
	return nil
}

// ReplyToMessageID sets the value for the reply_to_message_id parameter when making a call to any Telegram Bot API method.
func (params *telegramParameters) ReplyToMessageID(value int64) error {
	params.values.Set("reply_to_message_id", strconv.FormatInt(value, 10))
	return nil
}

// ReplyMarkup sets the value for the reply_markup parameter when making a call to any Telegram Bot API method.
func (params *telegramParameters) ReplyMarkup(value string) error {
	if len(value) == 0 {
		return errors.New("the reply_markup value cannot have zero length")
	}
	params.values.Set("reply_markup", value)
	return nil
}
