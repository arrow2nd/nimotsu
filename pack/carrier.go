package pack

import "sort"

// CarrierName : 配送業者名
type CarrierName string

type carrier struct {
	key      string
	tracking func(string) ([]status, error)
}

var carriers = map[CarrierName]*carrier{}

// GetCarriers : 配送業者のリストを取得
func GetCarriers() map[CarrierName]string {
	cs := map[CarrierName]string{}

	for name, c := range carriers {
		cs[name] = c.key
	}

	return cs
}

// GetCarrierNames : 配送業者名のリストを取得
func GetCarrierNames() []CarrierName {
	names := []CarrierName{}

	for name := range carriers {
		names = append(names, name)
	}

	sort.Slice(names, func(i, j int) bool { return names[i] < names[j] })

	return names
}
