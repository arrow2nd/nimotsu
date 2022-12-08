package pack

import "sort"

// Carrier : 配送業者名
type Carrier string

// CarrierInfo : 配送業者の情報
type CarrierInfo struct {
	Key    string
	Alias  string
	NameEn string
}

type carrier struct {
	*CarrierInfo
	tracking func(string) ([]status, error)
}

var carriers = map[Carrier]*carrier{}

// GetCarriers : 配送業者のリストを取得
func GetCarriers() map[Carrier]CarrierInfo {
	list := map[Carrier]CarrierInfo{}

	for name, c := range carriers {
		list[name] = *c.CarrierInfo
	}

	return list
}

// GetCarrierNames : 配送業者名のリストを取得
func GetCarrierNames() []Carrier {
	names := []Carrier{}

	for name := range carriers {
		names = append(names, name)
	}

	sort.Slice(names, func(i, j int) bool { return names[i] < names[j] })

	return names
}
