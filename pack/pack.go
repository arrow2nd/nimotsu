package pack

import "errors"

// Package : 荷物の情報
type Package struct {
	Carrier  CarrierName
	Number   string
	Comment  string
	statuses []status
}

// status : 配送状態
type status struct {
	date    string
	message string
	office  string
}

// Tracking : 追跡
func (p *Package) Tracking() error {
	c, ok := carriers[p.Carrier]
	if !ok {
		return errors.New("no carrier specified")
	}

	statuses, err := c.tracking(p.Number)
	if err != nil {
		return err
	}

	p.statuses = statuses
	return nil
}
