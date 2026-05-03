package main

import (
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
