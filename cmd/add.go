package cmd

import (
	"errors"
	"fmt"

	"github.com/arrow2nd/nimotsu/pack"
	"github.com/spf13/cobra"
)

func (c *Cmd) newAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:     "add [tracking number]",
		Short:   "Add the package",
		Long:    "Add the package to the list.",
		Example: "  nimotsu add --japanpost 112233445566 --comment \"DVD\"",
		Args:    cobra.ExactValidArgs(1),
		RunE:    c.execAddCmd,
	}

	setCarrierFlags(addCmd)
	addCmd.Flags().StringP("comment", "m", "", "add comment")

	return addCmd
}

func (c *Cmd) execAddCmd(cmd *cobra.Command, args []string) error {
	number := args[0]

	// フラグから配送業者を取得
	carrierName, err := getCarrierName(cmd.Flags())
	if err != nil {
		return nil
	}

	// リストに登録済みかチェック
	if c.list.Exists(number) {
		return errors.New("this tracking number exists in the list")
	}

	pkg := &pack.Package{
		Carrier: carrierName,
		Number:  number,
		Comment: "",
	}

	// 追跡番号が正しいかチェック
	if err := pkg.Tracking(); err != nil {
		return fmt.Errorf("this tracking number is wrong: %w", err)
	}

	comment, _ := cmd.Flags().GetString("comment")

	// コメントが無い場合入力
	if comment == "" {
		result, err := inputComment()
		if err != nil {
			return err
		}

		comment = result
	}

	c.list.AddItem(pkg)

	if err := c.list.Save(); err != nil {
		return err
	}

	showSuccessMessage("Added!")
	return nil
}
