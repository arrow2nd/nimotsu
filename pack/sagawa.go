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
	baseURL := "https://k2k.sagawa-exp.co.jp/p/web/okurijosearch.do"

	val := url.Values{}
	val.Add("okurijoNo", trackingNumber)

	doc, err := fetch(baseURL, val)
	if err != nil {
		return nil, err
	}

	results := []status{}
	table := doc.Find(".table_okurijo_detail2").Eq(1)

	table.Find("tr").Each(func(i int, s *goquery.Selection) {
		// 表のタイトルを読み飛ばす
		if i == 0 {
			return
		}

		td := s.Find("td").Map(func(_ int, s *goquery.Selection) string {
			return removeConsecutiveSpace(s.Text())
		})

		results = append(results, status{
			date:    td[1],
			message: td[0][3:], // 先頭の文字を削除
			office:  td[2],
		})
	})

	if len(results) == 0 {
		return nil, createNotFoundError(trackingNumber)
	}

	return results, nil
}
