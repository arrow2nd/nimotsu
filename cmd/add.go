package cmd

import (
	"fmt"

	"github.com/arrow2nd/nimotsu/list"
	"github.com/arrow2nd/nimotsu/pack"
	"github.com/spf13/cobra"
)

func (c *Cmd) newAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:     "add [tracking number]",
		Short:   "Add the package",
		Long:    "Add the package to the list.",
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
	number := args[0]

	// フラグから運送業者名を取得
	carrier, err := getCarrierName(cmd.Flags())
	if err != nil {
		return err
	}

	comment, _ := cmd.Flags().GetString("comment")
	if comment == "" {
		comment = "なし"
	}

	// 追跡番号が正しいかチェック
	pack := pack.New(carrier, number, "")
	if err = pack.Tracking(); err != nil {
		return fmt.Errorf("the tracking number is wrong")
	}

	c.list.AddItem(&list.Item{
		Carrier: carrier,
		Number:  number,
		Comment: comment,
	})
	c.list.Save()

	fmt.Println("Success: added (" + number + ")")
	return nil
}
