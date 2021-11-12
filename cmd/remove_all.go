package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (c *Cmd) newRemoveAllCmd() *cobra.Command {
	removeAllCmd := &cobra.Command{
		Use:     "all",
		Short:   "Remove all packages",
		Long:    "Remove all packages in the list.",
		Example: "  nimotsu remove all",
		Args:    cobra.NoArgs,
		Run:     c.execRemoveAllCmd,
	}

	return removeAllCmd
}

func (c *Cmd) execRemoveAllCmd(cmd *cobra.Command, args []string) {
	c.list.Clear()
	c.list.Save()

	fmt.Println("Success: all removed")
}
