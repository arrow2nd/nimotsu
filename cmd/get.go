package cmd

import (
	"fmt"

	"github.com/arrow2nd/nimotsu/track"
	"github.com/spf13/cobra"
)

var (
	isJapanPost bool
	isYamato    bool
	isSagawa    bool
)

var getCmd = &cobra.Command{
	Use:     "get [track number]",
	Short:   "track your package",
	Long:    `track your package for the carrier you specify.`,
	Example: "  nimotsu get --japanpost 112233445566",
	Args:    cobra.MinimumNArgs(1),
	RunE:    execGetCmd,
}

func init() {
	getCmd.Flags().BoolVarP(&isJapanPost, "japanpost", "j", false, "track Japan Post")
	getCmd.Flags().BoolVarP(&isYamato, "yamato", "y", false, "track Yamato Transport")
	getCmd.Flags().BoolVarP(&isSagawa, "sagawa", "s", false, "track Sagawa Express")

	rootCmd.AddCommand(getCmd)
}

func execGetCmd(cmd *cobra.Command, args []string) error {
	// 業者の指定が正しいかチェック
	enabledFlagCount := 0
	if isJapanPost {
		enabledFlagCount++
	}
	if isYamato {
		enabledFlagCount++
	}
	if isSagawa {
		enabledFlagCount++
	}

	if enabledFlagCount > 1 {
		return fmt.Errorf("expected exactly one of `--japanpost`, `--yamato`, or `--sagawa` to be true")
	}

	// 追跡
	pack := track.New(args[0], "なし")
	switch {
	case isJapanPost:
		pack.TrackByJapanPost()
	case isYamato:
		pack.TrackByYamato()
	case isSagawa:
		pack.TrackBySagawa()
	}

	// 表示
	err := pack.View()
	if err != nil {
		return err
	}

	return nil
}
