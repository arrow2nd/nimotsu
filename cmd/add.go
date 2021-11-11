package cmd

import (
	"fmt"

	"github.com/arrow2nd/nimotsu/list"
	"github.com/spf13/cobra"
)

func (c *Cmd) newAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:     "add [tracking number]",
		Short:   "Add the package",
		Long:    `Add the package to the list`,
		Example: "  nimotsu add --japanpost 112233445566 --comment DVD",
		Args:    cobra.ExactValidArgs(1),
		RunE:    c.execAddCmd,
	}

	addCmd.Flags().BoolP("japanpost", "j", false, "track Japan Post")
	addCmd.Flags().BoolP("yamato", "y", false, "track Yamato Transport")
	addCmd.Flags().BoolP("sagawa", "s", false, "track Sagawa Express")
	addCmd.Flags().StringP("comment", "m", "", "add comment")

	return addCmd
}

func (c *Cmd) execAddCmd(cmd *cobra.Command, args []string) error {
	carrier, err := getCarrierName(cmd.Flags())
	if err != nil {
		return err
	}

	comment, _ := cmd.Flags().GetString("comment")
	if comment == "" {
		comment = "なし"
	}

	newItem := list.Item{
		Carrier: carrier,
		Number:  args[0],
		Comment: comment,
	}

	c.list.AddItem(newItem)
	c.list.Save()

	c.list.View()
	fmt.Println("Success: added")

	return nil
}
