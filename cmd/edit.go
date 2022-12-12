package cmd

import "github.com/spf13/cobra"

func (c *Cmd) newEditCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "edit",
		Short: "Edit the package comments",
		Long:  "Edit the package comments in the list.",
		Args:  cobra.NoArgs,
		RunE:  c.execEditCmd,
	}
}

func (c *Cmd) execEditCmd(cmd *cobra.Command, args []string) error {
	item, err := c.selectTrackingNumber()
	if err != nil {
		return nil
	}

	comment, err := inputComment()
	if err != nil {
		return nil
	}

	if err := c.list.ChangeComment(item.Number, comment); err != nil {
		return err
	}

	if err := c.list.Save(); err != nil {
		return err
	}

	showSuccessMessage("Edited!")
	return nil
}
