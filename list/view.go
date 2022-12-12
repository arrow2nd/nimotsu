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

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"追跡番号", "コメント", "配送業者"})
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetAutoMergeCells(true)
	table.AppendBulk(data)
	table.Render()

	return nil
}
