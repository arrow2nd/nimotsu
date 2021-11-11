package pack

import "fmt"

// PackInfo 荷物情報
type PackInfo struct {
	carrier  string
	number   string
	comment  string
	statuses []status
}

// status 状態
type status struct {
	date    string
	message string
	office  string
}

// New 生成
func New(carrier, tracknumber, comment string) *PackInfo {
	return &PackInfo{
		carrier: carrier,
		number:  tracknumber,
		comment: comment,
	}
}

// Tracking 追跡
func (p *PackInfo) Tracking() error {
	switch p.carrier {
	case JapanPost:
		p.trackByJapanPost()
	case YamatoTransport:
		p.trackByYamato()
	case SagawaExpress:
		p.trackBySagawa()
	default:
		return fmt.Errorf("no carrier specified")
	}

	return nil
}
