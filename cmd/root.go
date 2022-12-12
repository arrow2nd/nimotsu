package cmd

import (
	"github.com/arrow2nd/nimotsu/list"
	"github.com/spf13/cobra"
)

const noCommentMessage = "なし"

// Cmd : 本体
type Cmd struct {
	root *cobra.Command
	list *list.List
}

// New : 生成
func New(l *list.List) *Cmd {
	cmd := &Cmd{
		list: l,
		root: &cobra.Command{
			Use:   "nimotsu",
			Short: "📦 CLI tool to tracking packages",
			Long:  "📦 CLI tool to tracking packages in japan",
		},
	}

	cmd.root.SilenceUsage = true
	cmd.root.AddCommand(
		cmd.newGetCmd(),
		cmd.newAddCmd(),
		cmd.newRemoveCmd(),
		cmd.newEditCmd(),
		cmd.newListCmd(),
		cmd.newBrowseCmd(),
		cmd.newVersionCmd(),
	)

	return cmd
}

// Execute : 実行
func (c *Cmd) Execute() error {
	return c.root.Execute()
}
