package main

import (
	"github.com/arrow2nd/nimotsu/cmd"
	"github.com/arrow2nd/nimotsu/list"
)

func main() {
	list := list.New()
	list.Load()

	cmd := cmd.New(list)
	cmd.Execute()
}
