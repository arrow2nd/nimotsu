package cmd

import (
	"github.com/arrow2nd/nimotsu/pack"
	"github.com/spf13/cobra"
)

func newGetCmd() *cobra.Command {
	getCmd := &cobra.Command{
		Use:     "get [track number]",
		Short:   "track your package",
		Long:    `track your package for the carrier you specify.`,
		Example: "  nimotsu get --japanpost 112233445566",
		Args:    cobra.ExactValidArgs(1),
		RunE:    execGetCmd,
	}

	getCmd.Flags().BoolP("japanpost", "j", false, "track Japan Post")
	getCmd.Flags().BoolP("yamato", "y", false, "track Yamato Transport")
	getCmd.Flags().BoolP("sagawa", "s", false, "track Sagawa Express")

	return getCmd
}

func execGetCmd(cmd *cobra.Command, args []string) error {
	carrier, err := getCarrierName(cmd.Flags())
	if err != nil {
		return err
	}

	pack := pack.New(carrier, args[0], "なし")
	pack.Tracking()

	err = pack.View()
	if err != nil {
		return err
	}

	return nil
}
