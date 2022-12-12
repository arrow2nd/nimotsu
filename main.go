package main

import (
	"os"

	"github.com/arrow2nd/nimotsu/cmd"
	"github.com/arrow2nd/nimotsu/list"
)

const (
	ExitCodeOK int = iota
	ExitCodeErrHomeDir
	ExitCodeErrLoad
	ExitCodeErrExec
)

func main() {
	list := list.New()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		os.Exit(ExitCodeErrHomeDir)
	}

	list.SetDir(homeDir)
	if err := list.Load(); err != nil {
		os.Exit(ExitCodeErrLoad)
	}

	cmd := cmd.New(list)
	if err := cmd.Execute(); err != nil {
		os.Exit(ExitCodeErrExec)
	}
}
