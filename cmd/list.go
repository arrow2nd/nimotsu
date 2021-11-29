package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (c *Cmd) newListCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Display the packages",
		Long:    "Display the packages in the list.",
		Example: "  nimotsu list",
		Args:    cobra.NoArgs,
		Run:     c.execListCmd,
	}

	return listCmd
}

func (c *Cmd) execListCmd(cmd *cobra.Command, args []string) {
	if err := c.list.View(); err != nil {
		fmt.Println("There's nothing!")
	}
}
