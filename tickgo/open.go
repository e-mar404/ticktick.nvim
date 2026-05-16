package main

import (
	"fmt"
	"net/url"
	"os/exec"
	"runtime"
)

func openURL(rawUrl string) error {
	var baseCmd string
	var msg string

	switch msg = runtime.GOOS; msg {
	case "darwin":
		baseCmd = "open"
	case "windows":
		baseCmd = "start"
	default:
		// this might not be best since not all other GOOS values have xdg-open but we'll leave it like this for now
		baseCmd = "xdg-open"
	}

	logger.Debug("opening url", "url", rawUrl, "GOOS", msg)
	return exec.Command(baseCmd, rawUrl).Run()
}
