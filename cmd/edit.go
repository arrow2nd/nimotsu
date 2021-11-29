package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func (c *Cmd) newEditCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "edit [tracking number] [comment]",
		Aliases: []string{"ed"},
		Short:   "Edit the package comments",
		Long:    "Edit the package comments in the list.",
		Example: "  nimotsu edit 112233445566 \"Blu-ray\"",
		Args:    cobra.ExactValidArgs(2),
		RunE:    c.execEditCmd,
	}
}

func (c *Cmd) execEditCmd(cmd *cobra.Command, args []string) error {
	number := args[0]
	comment := args[1]

	if len(comment) == 0 {
		comment = noCommentMessage
	}

	err := c.list.ChangeComment(number, comment)
	if err != nil {
		return err
	}

	if err := c.list.Save(); err != nil {
		return err
	}

	fmt.Printf("Edited: %s / %s\n", number, comment)
	return nil
}
