package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

var (
	RPC_PORT      string
	CALLBACK_PORT string
	verbose       bool
	logger        *log.Logger
	styles        = newStyles()
)

func init() {
	flag.StringVar(&RPC_PORT, "rpcPort", "", "port to start rpc server on")
	flag.StringVar(&CALLBACK_PORT, "callbackPort", "", "port to start callback server on")
	flag.BoolVar(&verbose, "v", false, "will output all log messages")
	flag.Parse()

	if RPC_PORT == "" || CALLBACK_PORT == "" {
		log.Error("Both flags rpcPort and callbackPort are required\n")
		os.Exit(1)
	}

	logger = log.New(os.Stderr)
	logger.SetReportCaller(true)
	logger.SetReportTimestamp(true)
	logger.SetStyles(styles.LoggerStyles)

	if verbose {
		logger.SetLevel(log.DebugLevel)
	}

	logger.Debug("ports set from flags",
		"rpc_port", RPC_PORT,
		"callback_port", CALLBACK_PORT,
	)
}

func main() {
	startRPCServer()
}
