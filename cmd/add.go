package cmd

import (
	"github.com/spf13/cobra"
)

func newAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:     "add [tracking number]",
		Short:   "add tracking number",
		Long:    ``,
		Example: "  nimotsu add --japanpost 112233445566 -m DVD",
		Args:    cobra.ExactValidArgs(1),
		RunE:    execAddCmd,
	}

	addCmd.Flags().BoolP("japanpost", "j", false, "track Japan Post")
	addCmd.Flags().BoolP("yamato", "y", false, "track Yamato Transport")
	addCmd.Flags().BoolP("sagawa", "s", false, "track Sagawa Express")
	addCmd.Flags().StringP("comment", "m", "", "add comment")

	return addCmd
}

func execAddCmd(cmd *cobra.Command, args []string) error {
	// carrier, err := getCarrierName(cmd.Flags())
	// if err != nil {
	// 	return err
	// }

	return nil
}
