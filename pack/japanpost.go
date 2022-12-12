package pack

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// CarrierJapanPost : 日本郵便
const CarrierJapanPost CarrierName = "日本郵便"

func init() {
	carriers[CarrierJapanPost] = &carrier{
		CarrierInfo: &CarrierInfo{
			Key:    "japanpost",
			Alias:  "j",
			NameEn: "Japan Post",
		},
		tracking: trackingByJapanPost,
	}
}

func trackingByJapanPost(trackingNumber string) ([]status, error) {
	const (
		trackingURL = "https://trackings.post.japanpost.jp/services/srv/search/direct"
		fieldMax    = 6
	)

	val := url.Values{}
	val.Add("searchKind", "S002")
	val.Add("locale", "ja")
	val.Add("reqCodeNo1", trackingNumber)

	doc, err := fetch(trackingURL, val)
	if err != nil {
		return nil, err
	}

	results := []status{}
	field := [fieldMax]string{}

	doc.Find("[summary='履歴情報'] td").Each(func(i int, s *goquery.Selection) {
		// 配達状況を追加
		if (i+1)%fieldMax == 0 {
			results = append(results, status{
				date:    field[0],
				message: field[1],
				office:  removeConsecutiveSpace(field[4] + " " + field[3]),
			})
		}

		field[i%fieldMax] = s.Text()
	})

	if len(results) == 0 {
		return nil, createNotFoundError(trackingNumber)
	}

	return results, nil
}
