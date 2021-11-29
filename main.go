package main

import (
	"github.com/arrow2nd/nimotsu/cmd"
	"github.com/arrow2nd/nimotsu/list"
)

func main() {
	list := list.New()
	if err := list.Load(); err != nil {
		panic(err)
	}

	cmd := cmd.New(list)
	cmd.Execute()
}
