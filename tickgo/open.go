package main

import (
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
	urlValues := url.Values{}
	urlValues.Add("client_id", clientID)
	urlValues.Add("scope", "tasks:write tasks:read")
	urlValues.Add("state", "state")                                 // TODO: needs actual random state
	urlValues.Add("redirect_uri", "http://127.0.0.1:9090/callback") // TODO: needs to be an env var
	urlValues.Add("response_type", "code")

	return openURL(TickTickOAuthURL + "/authorize?" + urlValues.Encode())
}
