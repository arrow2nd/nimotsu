package pack

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// CreateTableData 表データを作成
func (p *PackInfo) CreateViewData() [][]string {
	data := [][]string{}

	for _, status := range p.statuses {
		data = append(data, []string{
			p.number,
			p.comment,
			p.carrier,
			status.message,
			status.date,
			status.office,
		})
	}

	return data
}

// View 追跡状況を表示
func (p *PackInfo) View() error {
	data := p.CreateViewData()

	if len(data) == 0 {
		return fmt.Errorf("tracking number or shipping carrier is incorrect")
	}

	ShowTable(&data)

	return nil
}

// ShowTable テーブルを表示
func ShowTable(data *[][]string) {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"追跡番号", "コメント", "運送業者", "配達状況", "日時", "営業所"})
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetAutoMergeCells(true)
	table.AppendBulk(*data)
	table.Render()
}
