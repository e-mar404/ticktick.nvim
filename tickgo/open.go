package main

import (
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
)

func openURL(url string) error {
	var baseCmd string

	switch runtime.GOOS {
	case "darwin":
		baseCmd = "open"
	case "windows":
		baseCmd = "start"
	default:
		// this might not be best since not all other GOOS values have xdg-open but we'll leave it like this for now
		baseCmd = "xdg-open"
	}

	return exec.Command(baseCmd, url).Run()
}

func openOAuthPage(clientID string) error {
	redirectURI := fmt.Sprintf("http://127.0.0.1%s/callback", CALLBACK_PORT)
	urlValues := url.Values{}
	urlValues.Add("client_id", clientID)
	urlValues.Add("scope", "tasks:write tasks:read")
	urlValues.Add("state", "state") // TODO: needs actual random state and has to be checked when it comes back to make sure it is the same
	urlValues.Add("redirect_uri", redirectURI)
	urlValues.Add("response_type", "code")

	return openURL(TickTickOAuthURL + "/authorize?" + urlValues.Encode())
}
