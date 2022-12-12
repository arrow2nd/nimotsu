package cmd

import (
	"github.com/arrow2nd/nimotsu/pack"
	"github.com/spf13/cobra"
)

func (c *Cmd) newGetCmd() *cobra.Command {
	getCmd := &cobra.Command{
		Use:     "get [tracking number]",
		Short:   "Track Package",
		Long:    "Track package for the carrier you specify.",
		Example: "  nimotsu get --japanpost 112233445566",
		Args:    cobra.ExactValidArgs(1),
		RunE:    c.execGetCmd,
	}

	setCarrierFlags(getCmd)
	getCmd.AddCommand(c.newGetAllCmd())

	return getCmd
}

func (c *Cmd) execGetCmd(cmd *cobra.Command, args []string) error {
	carrierName, err := getCarrierName(cmd.Flags())
	if err != nil {
		return nil
	}

	pkg := pack.Package{
		Carrier: carrierName,
		Number:  args[0],
		Comment: noCommentMessage,
	}

	if err := pkg.Tracking(); err != nil {
		return err
	}

	return pkg.View()
}
