package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func (c *Cmd) newRemoveCmd() *cobra.Command {
	removeCmd := &cobra.Command{
		Use:     "remove",
		Aliases: []string{"rm"},
		Short:   "Remove the package",
		Long:    "Remove a package from the list.",
		Args:    cobra.NoArgs,
		RunE:    c.execRemoveCmd,
	}

	removeCmd.AddCommand(c.newRemoveAllCmd())

	return removeCmd
}

func (c *Cmd) execRemoveCmd(cmd *cobra.Command, args []string) error {
	if c.list.IsEmpty() {
		return errors.New("list is empty")
	}

	item, err := c.selectTrackingNumber()
	if err != nil {
		return nil
	}

	if err := c.list.RemoveItem(item.Number); err != nil {
		return err
	}

	if err := c.list.Save(); err != nil {
		return err
	}

	showSuccessMessage("Removed!")
	return nil
}
