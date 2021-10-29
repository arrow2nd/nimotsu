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
		if len(args) > 1 {
			fmt.Println("Too many arguments!")
			return
		}

		isJP, _ := cmd.Flags().GetBool("japanpost")
		isKuroneko, _ := cmd.Flags().GetBool("kuroneko")
		isSagawa, _ := cmd.Flags().GetBool("sagawa")

		fmt.Println(isJP, isKuroneko, isSagawa)
		fmt.Println(args[0])
	},
}

func init() {
	getCmd.Flags().BoolP("japanpost", "j", false, "Track Japan Post")
	getCmd.Flags().BoolP("kuroneko", "k", false, "Track Yamato Transport")
	getCmd.Flags().BoolP("sagawa", "s", false, "Tracking Sagawa Express")

	rootCmd.AddCommand(getCmd)
}
