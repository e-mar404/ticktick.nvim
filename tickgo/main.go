package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

var (
	RPC_PORT      string
	CALLBACK_PORT string
	verbose       bool
	l             *log.Logger
)

func init() {
	flag.StringVar(&RPC_PORT, "rpcPort", "", "port to start rpc server on")
	flag.StringVar(&CALLBACK_PORT, "callbackPort", "", "port to start callback server on")
	flag.BoolVar(&verbose, "v", false, "will output all log messages")
	flag.Parse()

	if RPC_PORT == "" || CALLBACK_PORT == "" {
		fmt.Printf("Both flags rpcPort and callbackPort are required\n")
		os.Exit(1)
	}

	l = log.New(os.Stderr)
	l.SetReportCaller(true)
	l.SetReportTimestamp(true)

	styles := log.DefaultStyles()

	styles.Levels[log.DebugLevel] = lipgloss.NewStyle().
		SetString(strings.ToUpper(log.DebugLevel.String())).
		Bold(true).
		MaxWidth(5).
		Foreground(lipgloss.Color("63"))

	styles.Levels[log.ErrorLevel] = lipgloss.NewStyle().
		SetString(strings.ToUpper(log.DebugLevel.String())).
		Bold(true).
		MaxWidth(5).
		Foreground(lipgloss.Color("63"))

	styles.Levels[log.FatalLevel] = lipgloss.NewStyle().
		SetString(strings.ToUpper(log.DebugLevel.String())).
		Bold(true).
		MaxWidth(5).
		Foreground(lipgloss.Color("63"))

	l.SetStyles(styles)

	if verbose {
		l.SetLevel(log.DebugLevel)
	}

	l.Debug("ports set from flags",
		"rpc_port", RPC_PORT,
		"callback_port", CALLBACK_PORT,
	)
}

func main() {
	startRPCServer()
}
