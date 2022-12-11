package cmd

import "github.com/spf13/cobra"

func (c *Cmd) newRemoveAllCmd() *cobra.Command {
	removeAllCmd := &cobra.Command{
		Use:     "all",
		Short:   "Remove all packages",
		Long:    "Remove all packages in the list.",
		Example: "  nimotsu remove all",
		Args:    cobra.NoArgs,
		RunE:    c.execRemoveAllCmd,
	}

	return removeAllCmd
}

func (c *Cmd) execRemoveAllCmd(cmd *cobra.Command, args []string) error {
	c.list.Clear()

	if err := c.list.Save(); err != nil {
		return err
	}

	showSuccessMessage("All removed!")
	return nil
}
