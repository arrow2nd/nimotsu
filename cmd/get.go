package cmd

import (
	"fmt"
	"os"

	"github.com/arrow2nd/nimotsu/track"
	"github.com/spf13/cobra"
)

var (
	isJapanPost bool
	isYamato    bool
	isSagawa    bool
)

var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "track your package",
	Long:    ``,
	Example: "  numotsu get [<tracking number>]",
	Run:     execGetCmd,
}

func init() {
	// 業者指定フラグ
	getCmd.Flags().BoolVarP(&isJapanPost, "japanpost", "j", false, "track Japan Post")
	getCmd.Flags().BoolVarP(&isYamato, "yamato", "y", false, "track Yamato Transport")
	getCmd.Flags().BoolVarP(&isSagawa, "sagawa", "s", false, "track Sagawa Express")

	rootCmd.AddCommand(getCmd)
}

func execGetCmd(cmd *cobra.Command, args []string) {
	// 引数エラー
	if len(args) > 1 {
		fmt.Fprintln(os.Stderr, "[Error] Too many arguments")
		return
	}

	// 追加済みの番号を全て追跡
	if len(args) == 0 {
		fmt.Println("Get All Added Number")
		return
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
	default:
		fmt.Fprintln(os.Stderr, "[Error] Please specify the shipping carrier")
		return
	}

	// 表示
	data := pack.CreateTableData()
	if len(data) == 0 {
		fmt.Fprintln(os.Stderr, "[Error] The tracking number or shipping carrier is incorrect")
		return
	}

	track.ShowTable(&data)
}
