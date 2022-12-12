package cmd

import (
	"fmt"
	"io"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

func (c *Cmd) newBrowseCmd() *cobra.Command {
	browseCmd := &cobra.Command{
		Use:     "browse",
		Aliases: []string{"br"},
		Short:   "Open tracking page in browser",
		Long:    "Open tracking page in browser.",
		Args:    cobra.NoArgs,
		RunE:    c.execBrowseCmd,
	}

	return browseCmd
}

func (c *Cmd) execBrowseCmd(cmd *cobra.Command, args []string) error {
	if c.list.IsEmpty() {
		return fmt.Errorf("list is empty")
	}

	number, err := c.selectTrackingNumber()
	if err != nil {
		return nil
	}

	url, err := c.list.GetTrackingURL(number)
	if err != nil {
		return err
	}

	browser.Stdout = io.Discard
	browser.Stderr = io.Discard

	return browser.OpenURL(url)
}
