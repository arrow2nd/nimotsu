package pack

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// CarrierSagawa : 佐川急便
const CarrierSagawa CarrierName = "佐川急便"

func init() {
	carriers[CarrierSagawa] = &carrier{
		CarrierInfo: &CarrierInfo{
			Key:    "sagawa",
			Alias:  "s",
			NameEn: "Sagawa Express",
		},
		tracking: trackingBySagawa,
	}
}

func trackingBySagawa(trackingNumber string) ([]status, error) {
	const trackingURL = "https://k2k.sagawa-exp.co.jp/p/web/okurijosearch.do"

	val := url.Values{}
	val.Add("okurijoNo", trackingNumber)

	doc, err := fetch(trackingURL, val)
	if err != nil {
		return nil, err
	}

	results := []status{}

	// TODO: Eqで2個目の要素をもらってインデントを減らしたい
	doc.Find("table.table_basic.table_okurijo_detail2").Each(func(i int, s *goquery.Selection) {
		// 1ループ目は荷物情報なので読み飛ばす
		if i == 0 {
			return
		}

		s.Find("tr").Each(func(i int, s *goquery.Selection) {
			// 表のタイトルを読み飛ばす
			if i == 0 {
				return
			}

			field := []string{}
			// TODO: Map()を使う形にしたい
			s.Find("td").Each(func(i int, s *goquery.Selection) {
				field[i] = removeConsecutiveSpace(s.Text())
			})

			results = append(results, status{
				date:    field[1],
				message: field[0][3:], // 先頭の文字を削除
				office:  field[2],
			})
		})
	})

	if len(results) == 0 {
		return nil, createNotFoundError(trackingNumber)
	}

	return results, nil
}
