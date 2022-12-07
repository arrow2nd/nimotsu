package pack

import "errors"

// Package : 荷物の情報
type Package struct {
	carrierName    CarrierName
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
func New(name CarrierName, number, comment string) *Package {
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
