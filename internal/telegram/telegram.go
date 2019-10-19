package main

import (
	"bytes"
	"emersyx.net/emersyx/api/tgapi"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

type apiResponse struct {
	OK     bool            `json:"ok"`
	Result json.RawMessage `json:"result"`
}

// telegramURL builds the URL for the Telegram Bot API, with the API token of the telegramGateway instance and for the
// specified API method.
func (gw *telegramGateway) telegramURL(apiMethod string) string {
	return "https://api.telegram.org/bot" + gw.apiToken + "/" + apiMethod
}

// sendAPIRequest perfoms an HTTP GET or POST request to the Telegram Bot API. If parameters are given as argument, then
// a POST request is made, otherwise a GET request is made. The response is parsed and returned.
func (gw *telegramGateway) sendAPIRequest(apiMethod string, params url.Values) (*apiResponse, error) {
	var err error
	var resp *http.Response
	apiresp := new(apiResponse)

	if params == nil {
		resp, err = http.Get(gw.telegramURL(apiMethod))
	} else {
		resp, err = http.PostForm(gw.telegramURL(apiMethod), params)
	}

	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("api response - %s\n%s", resp.Status, buf.String())
	}

	err = json.Unmarshal(buf.Bytes(), apiresp)
	if err != nil {
		return nil, err
	}

	if apiresp.OK == false {
		return nil, errors.New("the ok field in the Bot API response is false")
	}

	return apiresp, nil
}

// setAPIToken sets the Telegram Bot API token for the telegramGateway instance and validates it by performing a request
// with the getMe method.
func (gw *telegramGateway) setAPIToken(token string) error {
	// check that the token is not an empty string
	if len(token) == 0 {
		return errors.New("the API token cannot have zero length")
	}
	gw.apiToken = token
	// perform a call to the getMe method to test the token
	if _, err := gw.GetMe(); err != nil {
		return err
	}
	return nil
}

// NewTelegramParameters creates and returns a new telegramParameters object. This object can then be used to configure
// parameters when performing calls to the Telegram Bot API (e.g. see the telegramGateway.SendMessage method).
func (gw *telegramGateway) NewTelegramParameters() tgapi.TelegramParameters {
	params := new(telegramParameters)
	params.values = make(map[string][]string)
	return params
}

// getUpdates performs calls to the getUpdates method of the Telegram Bot API and converts the data into tgapi.Update
// instances.
func (gw *telegramGateway) getUpdates(offset int64) (updates []tgapi.Update, err error) {
	params := gw.NewTelegramParameters().(*telegramParameters)
	params.Offset(offset)
	params.Limit(gw.updatesLimit)
	params.Timeout(gw.updatesTimeout)

	apiresp, err := gw.sendAPIRequest("getUpdates", params.values)
	if err != nil {
		return updates, err
	}

	err = json.Unmarshal(apiresp.Result, &updates)
	if err != nil {
		return updates, err
	}

	return updates, nil
}

// GetMe performs a call to the getMe method of the Telegram Bot API.
func (gw *telegramGateway) GetMe() (user tgapi.User, err error) {
	apiresp, err := gw.sendAPIRequest("getMe", nil)
	if err != nil {
		return user, err
	}

	err = json.Unmarshal(apiresp.Result, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}

// SendMessage performs a call to the sendMessage method of the Telegram Bot API.
func (gw *telegramGateway) SendMessage(params tgapi.TelegramParameters) (msg tgapi.Message, err error) {
	cparams, err := paramVals(params)
	if err != nil {
		return msg, err
	}

	apiresp, err := gw.sendAPIRequest("sendMessage", cparams)
	if err != nil {
		return msg, err
	}

	err = json.Unmarshal(apiresp.Result, &msg)
	if err != nil {
		return msg, err
	}

	return msg, nil
}
