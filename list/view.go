package list

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func (l *List) createTableData() [][]string {
	data := [][]string{}

	for _, item := range l.items {
		data = append(data, []string{
			item.Number,
			item.Comment,
			item.Carrier,
		})
	}

	return data
}

// View リストを表示する
func (l *List) View() {
	table := tablewriter.NewWriter(os.Stdout)
	data := l.createTableData()

	table.SetHeader([]string{"追跡番号", "コメント", "運送業者"})
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetAutoMergeCells(true)
	table.AppendBulk(data)
	table.Render()
}
