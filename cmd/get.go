package cmd

import (
	"fmt"
	"log"

	"github.com/arrow2nd/nimotsu/track"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "track your package",
	Long:    ``,
	Example: "  numotsu get [<tracking number>]",
	Run: func(cmd *cobra.Command, args []string) {
		// 追加済みの番号を全て追跡
		if len(args) == 0 {
			fmt.Println("Get All Added Number")
			return
		}

		// 引数エラー
		if len(args) > 1 {
			log.Fatalln("[Error] Too many arguments")
		}

		pack := track.New(args[0], "なし")

		isJP, _ := cmd.Flags().GetBool("japanpost")
		isYamato, _ := cmd.Flags().GetBool("yamato")
		isSagawa, _ := cmd.Flags().GetBool("sagawa")

		// 業者ごとに追跡
		switch {
		case isJP:
			pack.TrackByJapanPost()
		case isYamato:
			pack.TrackByYamato()
		case isSagawa:
			pack.TrackBySagawa()
		default:
			log.Fatalln("[Error] Please specify the shipping carrier")
		}

		data := pack.CreateTableData()

		track.ShowTable(&data)
	},
}

func init() {
	// 業者指定フラグ
	getCmd.Flags().BoolP("japanpost", "j", false, "Track Japan Post")
	getCmd.Flags().BoolP("yamato", "y", false, "Track Yamato Transport")
	getCmd.Flags().BoolP("sagawa", "s", false, "Tracking Sagawa Express")

	rootCmd.AddCommand(getCmd)
}
