package cmd

import (
	"github.com/arrow2nd/nimotsu/list"
	"github.com/spf13/cobra"
)

type Cmd struct {
	root *cobra.Command
	list *list.List
}

// New 生成
func New(l *list.List) *Cmd {
	cmd := &Cmd{
		list: l,
	}

	cmd.root = &cobra.Command{
		Use:   "nimotsu",
		Short: "CLI tool to track packages 📦",
		Long:  ``, // TODO: あとで書く
	}
	cmd.root.AddCommand(cmd.newGetCmd(), cmd.newAddCmd())

	return cmd
}

// Execute 実行
func (c *Cmd) Execute() {
	c.root.Execute()
}
