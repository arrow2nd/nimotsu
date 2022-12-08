package pack

import (
	"fmt"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// CarrierSagawa : 佐川急便
const CarrierSagawa Carrier = "佐川急便"

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
	const (
		trackingURL = "https://k2k.sagawa-exp.co.jp/p/web/okurijosearch.do"
		fieldMax    = 3
	)

	val := url.Values{}
	val.Add("okurijoNo", trackingNumber)

	doc, err := fetchBody(trackingURL, val)
	if err != nil {
		return nil, err
	}

	var results []status

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

			var field [fieldMax]string
			s.Find("td").Each(func(i int, s *goquery.Selection) {
				field[i] = removeConsecutiveSpace(s.Text())
			})

			results = append(results, status{
				date:    fmt.Sprintf("%d/%s", time.Now().Year(), field[1]),
				message: field[0][3:], // 先頭の文字を削除
				office:  field[2],
			})
		})
	})

	if len(results) == 0 {
		return nil, fmt.Errorf("couldn't find the package (" + trackingNumber + ")")
	}

	return results, nil
}
