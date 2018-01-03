package tgbotapi

import(
    "errors"
    "fmt"
    "net/http"
    "net/url"
    "strconv"
)

const(
    API_URL = "https://api.telegram.org/"
)

var apiToken string

func Initialize(token string) error {
    // check that the token is not an empty string
    if len(token) == 0 {
        return errors.New("The API token cannot have zero length.")
    }

    // perform a call to the getMe method to test the token
    resp, err := http.Get(API_URL + "bot" + token + "/getMe")

    // check if the request was made and the returned status code
    if err != nil {
        return err
    } else if resp.StatusCode != http.StatusOK {
        // if the token is valid, then 200 OK is returned
        // if the token is not valid, then 401 Unauthorized is returned
        return errors.New("Invalid Telegram Bot authentication token.")
    }

    // save the api token
    apiToken = token

    return nil
}

// Use this method to receive incoming updates using long polling. An Array of Update objects is returned.
func GetUpdates(params url.Values) (string, error) {
    limit := params.Get("limit")
    if limit != "" {
        ilimit, err := strconv.ParseInt(limit, 10, 64)
        if err != nil {
            return "", err
        }
        if ilimit < 1 || ilimit > 100 {
            return "", errors.New("The value for the limit parameter must be between 1 and 100.")
        }
    }

    timeout := params.Get("timeout")
    if limit != "" {
        itimeout, err := strconv.ParseInt(timeout, 10, 64)
        if err != nil {
            return "", err
        }
        if itimeout < 0 {
            return "", errors.New("The value for the timeout parameter must not be negative.")
        }
    }

    resp, err := http.PostForm(tgurl("getUpdates"), params)
    if err != nil {
        return "", err
    } else if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("Returned HTTP status code is %s.", resp.Status)
    }

    return responseBody(resp), nil
}

// A simple method for testing your bot's auth token. Requires no parameters. Returns basic information about the bot in
// form of a User object.
func GetMe() (string, error) {
    resp, err := http.Get(tgurl("getMe"))
    if err != nil {
        return "", err
    } else if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("Returned HTTP status code is %s.", resp.Status)
    }

    return responseBody(resp), nil
}

// Use this method to send text messages. On success, the sent Message is returned.
func SendMessage(params url.Values) (string, error) {
    if checkParams(params, "chat_id", "text") == false {
        return "", errors.New("Mandatory parameters have not been set.")
    }

    pm := params.Get("parse_mode")
    if pm != "" && pm != "Markdown" && pm != "HTML" {
        return "", errors.New("The parse_mode value can only be set to Markdown or HTML.")
    }

    resp, err := http.PostForm(tgurl("sendMessage"), params)
    if err != nil {
        return "", err
    } else if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("Returned HTTP status code is %s.", resp.Status)
    }

    return responseBody(resp), nil
}
