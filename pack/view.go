package pack

import (
	"errors"
	"os"

	"github.com/olekukonko/tablewriter"
)

// CreateViewData : 表データを作成
func (p *Package) CreateViewData() [][]string {
	data := [][]string{}

	for _, status := range p.statuses {
		data = append(data, []string{
			p.Number,
			p.Comment,
			string(p.Carrier),
			status.message,
			status.date,
			status.office,
		})
	}

	return data
}

// View : 追跡状況を表示
func (p *Package) View() error {
	data := p.CreateViewData()
	if len(data) == 0 {
		return errors.New("tracking number or shipping carrier is incorrect")
	}

	ShowTable(&data)
	return nil
}

// ShowTable : テーブルを表示
func ShowTable(data *[][]string) {
	table := tablewriter.NewTable(os.Stdout)
	table.Header([]string{"追跡番号", "コメント", "配送業者", "配達状況", "日時", "営業所"})
	table.Bulk(*data)
	table.Render()
}
