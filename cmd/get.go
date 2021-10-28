package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "track your package",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(args[0])
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
