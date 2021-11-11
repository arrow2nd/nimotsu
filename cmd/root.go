package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nimotsu",
	Short: "CLI tool to track packages 📦",
	Long:  ``, // TODO: あとで書く
}

func init() {
	rootCmd.AddCommand(newGetCmd())
	rootCmd.AddCommand(newAddCmd())
}

func Execute() {
	rootCmd.Execute()
}
