package main

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

type Styles struct {
	LoggerStyles *log.Styles
}

func newStyles() Styles {
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

	return Styles{
		LoggerStyles: styles,
	}
}
