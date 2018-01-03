package main

import(
    "errors"
    "net/url"
    "strconv"
    "emersyx.net/emersyx_apis/emtgapi"
)

type TelegramParameters struct {
    values url.Values
}

func (params *TelegramParameters) Offset(value int64) error {
    params.values.Set("offset", strconv.FormatInt(value, 10))
    return nil
}

func (params *TelegramParameters) Limit(value uint) error {
    params.values.Set("limit", strconv.FormatUint(uint64(value), 10))
    return nil
}

func (params *TelegramParameters) Timeout(value uint) error {
    params.values.Set("timeout", strconv.FormatUint(uint64(value), 10))
    return nil
}

func (params *TelegramParameters) AllowedUpdates(values ...string) error {
    for _, value := range values {
        if len(value) > 0 {
            params.values.Add("allowed_updates", value)
        }
    }
    return nil
}

func (params *TelegramParameters) ChatID(value string) error {
    if len(value) == 0 {
        return errors.New("The chat_id value cannot have zero length.")
    }
    params.values.Set("chat_id", value)
    return nil
}

func (params *TelegramParameters) Text(value string) error {
    if len(value) == 0 {
        return errors.New("The text value cannot have zero length.")
    }
    params.values.Set("text", value)
    return nil
}

func (params *TelegramParameters) ParseMode(value string) error {
    if len(value) == 0 {
        return errors.New("The parse_mode value cannot have zero length.")
    }
    params.values.Set("parse_mode", value)
    return nil
}

func (params *TelegramParameters) DisableWebPagePreview(value bool) error {
    params.values.Set("disable_web_page_preview", strconv.FormatBool(value))
    return nil
}

func (params *TelegramParameters) DisableNotification(value bool) error {
    params.values.Set("disable_notification", strconv.FormatBool(value))
    return nil
}

func (params *TelegramParameters) ReplyToMessageID(value int64) error {
    params.values.Set("reply_to_message_id", strconv.FormatInt(value, 10))
    return nil
}

func (params *TelegramParameters) ReplyMarkup(value string) error {
    if len(value) == 0 {
        return errors.New("The reply_markup value cannot have zero length.")
    }
    params.values.Set("reply_markup", value)
    return nil
}

func NewTelegramParameters() emtgapi.TelegramParameters {
    params := new(TelegramParameters)
    params.values = make( map[string][]string )
    return params
}
