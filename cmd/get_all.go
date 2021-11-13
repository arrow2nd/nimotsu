package cmd

import (
	"sync"

	"github.com/arrow2nd/nimotsu/pack"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func (c *Cmd) newGetAllCmd() *cobra.Command {
	getAllCmd := &cobra.Command{
		Use:     "all",
		Short:   "Track all packages",
		Long:    "Track all packages in the list.",
		Example: "  nimotsu get all",
		Args:    cobra.NoArgs,
		RunE:    c.execGetAllCmd,
	}

	return getAllCmd
}

func (c *Cmd) execGetAllCmd(cmd *cobra.Command, args []string) error {
	results := [][]string{}
	eg := errgroup.Group{}
	mutex := sync.Mutex{}

	for _, item := range c.list.Get() {
		item := item

		eg.Go(func() error {
			pack := pack.New(item.Carrier, item.Number, item.Comment)

			if err := pack.Tracking(); err != nil {
				return err
			}

			mutex.Lock()
			results = append(results, pack.CreateViewData()...)
			mutex.Unlock()

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return err
	}

	pack.ShowTable(&results)
	return nil
}
