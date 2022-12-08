package pack

import "errors"

const dateFormat = "2006/01/02 15:04"

// Package : 荷物の情報
type Package struct {
	carrierName    Carrier
	trackingNumber string
	comment        string
	statuses       []status
}

// status : 配送状態
type status struct {
	date    string
	message string
	office  string
}

// New : 生成
func New(name Carrier, number, comment string) *Package {
	return &Package{
		carrierName:    name,
		trackingNumber: number,
		comment:        comment,
	}
}

// Tracking : 追跡
func (p *Package) Tracking() error {
	c, ok := carriers[p.carrierName]
	if !ok {
		return errors.New("no carrier specified")
	}

	statuses, err := c.tracking(p.trackingNumber)
	if err != nil {
		return err
	}

	p.statuses = statuses
	return nil
}
