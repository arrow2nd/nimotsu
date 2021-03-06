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
func New(carrier, tNumber, comment string) *PackInfo {
	return &PackInfo{
		carrier: carrier,
		number:  tNumber,
		comment: comment,
	}
}

// Tracking 追跡
func (p *PackInfo) Tracking() error {
	var err error

	switch p.carrier {
	case JapanPost:
		err = p.trackByJapanPost()
	case YamatoTransport:
		err = p.trackByYamato()
	case SagawaExpress:
		err = p.trackBySagawa()
	default:
		return fmt.Errorf("no carrier specified")
	}

	return err
}
