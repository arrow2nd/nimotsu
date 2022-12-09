package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "develop"

func (c *Cmd) newVersionCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"ver"},
		Short:   "Display current version",
		Long:    "Display current version of the CLI.",
		Example: "  nimotsu version",
		Args:    cobra.NoArgs,
		Run:     c.execVersionCmd,
	}

	return listCmd
}

func (c *Cmd) execVersionCmd(cmd *cobra.Command, args []string) {
	fmt.Printf("📦 nimotsu ver.%s\n", version)
}
