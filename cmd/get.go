package cmd

import (
	"fmt"

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
			fmt.Println("Error : Too many arguments!")
			return
		}

		isJP, _ := cmd.Flags().GetBool("japanpost")
		isYamato, _ := cmd.Flags().GetBool("yamato")
		isSagawa, _ := cmd.Flags().GetBool("sagawa")
		number := args[0]

		// 業者毎に分岐
		switch {
		case isJP:
			fmt.Println("JapanPost : " + number)
		case isYamato:
			fmt.Println("Yamato : " + number)
		case isSagawa:
			fmt.Println("Sagawa : " + number)
		default:
			// 追加済みの番号を検索
			fmt.Println("Search AddedList : " + number)
		}
	},
}

func init() {
	// 業者指定フラグ
	getCmd.Flags().BoolP("japanpost", "j", false, "Track Japan Post")
	getCmd.Flags().BoolP("yamato", "y", false, "Track Yamato Transport")
	getCmd.Flags().BoolP("sagawa", "s", false, "Tracking Sagawa Express")

	rootCmd.AddCommand(getCmd)
}
