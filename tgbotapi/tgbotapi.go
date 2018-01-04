package tgbotapi

import(
    "errors"
    "fmt"
    "net/http"
    "net/url"
    "strconv"
)

var apiToken string

// Initialize sets up the tgbotapi package and verifies that the provided Telegram Bot API token is valid by performing
// a call to the getMe method.
func Initialize(token string) error {
    // check that the token is not an empty string
    if len(token) == 0 {
        return errors.New("the API token cannot have zero length")
    }

    // perform a call to the getMe method to test the token
    resp, err := http.Get("https://api.telegram.org/bot" + token + "/getMe")

    // check if the request was made and the returned status code
    if err != nil {
        return err
    } else if resp.StatusCode != http.StatusOK {
        // if the token is valid, then 200 OK is returned
        // if the token is not valid, then 401 Unauthorized is returned
        return errors.New("invalid Telegram Bot authentication token")
    }

    // save the api token
    apiToken = token

    return nil
}

// GetUpdates perfoms a call to the getUpdates method of the Telegram Bot API.
func GetUpdates(params url.Values) (string, error) {
    limit := params.Get("limit")
    if limit != "" {
        ilimit, err := strconv.ParseInt(limit, 10, 64)
        if err != nil {
            return "", err
        }
        if ilimit < 1 || ilimit > 100 {
            return "", errors.New("the value for the limit parameter must be between 1 and 100")
        }
    }

    timeout := params.Get("timeout")
    if limit != "" {
        itimeout, err := strconv.ParseInt(timeout, 10, 64)
        if err != nil {
            return "", err
        }
        if itimeout < 0 {
            return "", errors.New("the value for the timeout parameter must not be negative")
        }
    }

    resp, err := http.PostForm(tgurl("getUpdates"), params)
    if err != nil {
        return "", err
    } else if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("returned HTTP status code is %s", resp.Status)
    }

    return responseBody(resp), nil
}

// GetMe performs a call to the getMe method of the Telegram Bot API.
func GetMe() (string, error) {
    resp, err := http.Get(tgurl("getMe"))
    if err != nil {
        return "", err
    } else if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("returned HTTP status code is %s", resp.Status)
    }

    return responseBody(resp), nil
}

// SendMessage performs a call to the sendMessage method of the Telegram Bot API.
func SendMessage(params url.Values) (string, error) {
    if checkParams(params, "chat_id", "text") == false {
        return "", errors.New("mandatory parameters have not been set")
    }

    pm := params.Get("parse_mode")
    if pm != "" && pm != "Markdown" && pm != "HTML" {
        return "", errors.New("the parse_mode value can only be set to Markdown or HTML")
    }

    resp, err := http.PostForm(tgurl("sendMessage"), params)
    if err != nil {
        return "", err
    } else if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("returned HTTP status code is %s", resp.Status)
    }

    return responseBody(resp), nil
}
