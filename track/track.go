package track

// Track 追跡情報
type Track struct {
	number   string
	statuses []Status
}

// Status 状態
type Status struct {
	Message string
	Date    string
	Memo    string
}

// New 生成
func New(tracknumber string) *Track {
	return &Track{
		number: tracknumber,
	}
}
