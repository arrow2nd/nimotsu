package cmd

import (
	"github.com/spf13/cobra"
)

func (c *Cmd) newGetAllCmd() *cobra.Command {
	getAllCmd := &cobra.Command{
		Use:     "all",
		Short:   "Track all packages",
		Long:    "Track all packages in the list.",
		Example: "  nimotsu get all",
		Args:    cobra.NoArgs,
		RunE:    c.execGetAllCmd,
	}

	return getAllCmd
}

func (c *Cmd) execGetAllCmd(cmd *cobra.Command, args []string) error {

	return nil
}
