package cmd

import (
	"github.com/arrow2nd/nimotsu/list"
	"github.com/spf13/cobra"
)

const version = "1.0.0"

type Cmd struct {
	root *cobra.Command
	list *list.List
}

// New 生成
func New(l *list.List) *Cmd {
	cmd := &Cmd{
		list: l,
		root: &cobra.Command{
			Use:   "nimotsu",
			Short: "CLI tool to track packages 📦",
			Long:  "CLI tool to track packages by tracking number 📦",
		},
	}

	cmd.root.SilenceUsage = true
	cmd.root.AddCommand(
		cmd.newGetCmd(),
		cmd.newAddCmd(),
		cmd.newRemoveCmd(),
		cmd.newListCmd(),
		cmd.newVersionCmd(),
	)

	return cmd
}

// Execute 実行
func (c *Cmd) Execute() {
	c.root.Execute()
}
