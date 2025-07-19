package list

import (
	"errors"
	"os"

	"github.com/olekukonko/tablewriter"
)

func (l *List) createTableData() [][]string {
	data := [][]string{}

	for _, pkg := range l.packages {
		data = append(data, []string{
			pkg.Number,
			pkg.Comment,
			string(pkg.Carrier),
		})
	}

	return data
}

// View : リストを表示する
func (l *List) View() error {
	data := l.createTableData()
	if len(data) == 0 {
		return errors.New("list is empty")
	}

	table := tablewriter.NewTable(os.Stdout)
	table.Header([]string{"追跡番号", "コメント", "配送業者"})
	table.Bulk(data)
	table.Render()

	return nil
}
