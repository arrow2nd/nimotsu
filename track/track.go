package track

// Track 追跡情報
type Track struct {
	company  string
	number   string
	memo     string
	statuses []Status
}

// Status 状態
type Status struct {
	Date    string
	Message string
	Office  string
}

// New 生成
func New(tracknumber, memo string) *Track {
	return &Track{
		number: tracknumber,
		memo:   memo,
	}
}

// CreateTableData 表データを作成
func (t *Track) CreateTableData() [][]string {
	data := [][]string{}

	for _, status := range t.statuses {
		data = append(data, []string{
			t.company,
			t.number,
			t.memo,
			status.Message,
			status.Date,
			status.Office,
		})
	}

	return data
}