package cmd

import (
	"github.com/arrow2nd/nimotsu/pack"
	"github.com/spf13/cobra"
)

func (c *Cmd) newGetCmd() *cobra.Command {
	getCmd := &cobra.Command{
		Use:     "get [track number]",
		Short:   "Track Package",
		Long:    "Track package for the carrier you specify.",
		Example: "  nimotsu get --japanpost 112233445566",
		Args:    cobra.ExactValidArgs(1),
		RunE:    c.execGetCmd,
	}

	getCmd.Flags().BoolP("japanpost", "j", false, "track Japan Post")
	getCmd.Flags().BoolP("yamato", "y", false, "track Yamato Transport")
	getCmd.Flags().BoolP("sagawa", "s", false, "track Sagawa Express")
	getCmd.AddCommand(c.newGetAllCmd())

	return getCmd
}

func (c *Cmd) execGetCmd(cmd *cobra.Command, args []string) error {
	carrier, err := getCarrierName(cmd.Flags())
	if err != nil {
		return err
	}

	pack := pack.New(carrier, args[0], "なし")
	err = pack.Tracking()
	if err != nil {
		return err
	}

	return pack.View()
}
