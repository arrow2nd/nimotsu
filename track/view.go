package track

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

// ShowTable 情報を表示
func ShowTable(data *[][]string) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"運送業者", "追跡番号", "メモ", "配達状況", "日時", "営業所"})
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetAutoMergeCells(true)
	table.AppendBulk(*data)
	table.Render()
}
