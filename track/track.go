package track

// Track 追跡情報
type Track struct {
	number   string
	statuses []Status
}

// Status 状態
type Status struct {
	Date    string
	Message string
	Office  string
}

// New 生成
func New(tracknumber string) *Track {
	return &Track{
		number: tracknumber,
	}
}
