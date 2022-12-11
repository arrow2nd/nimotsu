package main

import (
	"os"

	"github.com/arrow2nd/nimotsu/cmd"
	"github.com/arrow2nd/nimotsu/list"
)

func main() {
	list := list.New()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	list.SetDir(homeDir)
	if err := list.Load(); err != nil {
		panic(err)
	}

	cmd := cmd.New(list)
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
