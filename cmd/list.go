package cmd

import (
	"github.com/spf13/cobra"
)

func (c *Cmd) newListCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "",
		Long:    "",
		Example: "  nimotsu list",
		Args:    cobra.NoArgs,
		Run:     c.execListCmd,
	}

	return listCmd
}

func (c *Cmd) execListCmd(cmd *cobra.Command, args []string) {
	c.list.View()
}
