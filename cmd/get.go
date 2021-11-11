package cmd

import (
	"fmt"

	"github.com/arrow2nd/nimotsu/track"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get [track number]",
	Short:   "track your package",
	Long:    `track your package for the carrier you specify.`,
	Example: "  nimotsu get --japanpost 112233445566",
	Args:    cobra.ExactValidArgs(1),
	RunE:    execGetCmd,
}

func init() {
	getCmd.Flags().BoolP("japanpost", "j", false, "track Japan Post")
	getCmd.Flags().BoolP("yamato", "y", false, "track Yamato Transport")
	getCmd.Flags().BoolP("sagawa", "s", false, "track Sagawa Express")

	rootCmd.AddCommand(getCmd)
}

func execGetCmd(cmd *cobra.Command, args []string) error {
	carrier := carrier{}
	carrier.isJapanPost, _ = cmd.Flags().GetBool("japanpost")
	carrier.isYamato, _ = cmd.Flags().GetBool("yamato")
	carrier.isSagawa, _ = cmd.Flags().GetBool("sagawa")

	if !carrier.isValid() {
		return fmt.Errorf("expected exactly one of `--japanpost`, `--yamato`, or `--sagawa` to be true")
	}

	// 追跡
	pack := track.New(args[0], "なし")
	switch {
	case carrier.isJapanPost:
		pack.TrackByJapanPost()
	case carrier.isYamato:
		pack.TrackByYamato()
	case carrier.isSagawa:
		pack.TrackBySagawa()
	}

	// 表示
	err := pack.View()
	if err != nil {
		return err
	}

	return nil
}
