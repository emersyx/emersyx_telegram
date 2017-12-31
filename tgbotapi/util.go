package tgbotapi

import(
    "bytes"
    "net/http"
    "net/url"
)

// This function builds the URL for the Telegram Bot API for the API token used in the Initialize function and the
// specified method.
func tgurl(apiMethod string) string {
    return API_URL + "bot" + apiToken + "/" + apiMethod
}

// This function takes an *http.Response argument and returns the body of the response as a string.
func responseBody(resp *http.Response) string {
    buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    return buf.String()
}

// This function checks if all mandatory parameters (specified as argument) are present and not an empty string.
func checkParams(params url.Values, mandatory ...string) bool {
    for _, m := range mandatory {
        if params.Get(m) == "" {
            return false
        }
    }
    return true
}
