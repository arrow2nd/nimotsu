package pack

import (
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// CarrierYamato : ヤマト運輸
const CarrierYamato CarrierName = "ヤマト運輸"

func init() {
	carriers[CarrierYamato] = &carrier{
		CarrierInfo: &CarrierInfo{
			Key:    "yamato",
			Alias:  "y",
			NameEn: "Yamato Transport",
		},
		tracking: trackingByYamato,
	}
}

func trackingByYamato(trackingNumber string) ([]status, error) {
	const trackingURL = "https://toi.kuronekoyamato.co.jp/cgi-bin/tneko"

	val := url.Values{}
	val.Add("number00", "1") // 取得件数？
	val.Add("number01", trackingNumber)

	doc, err := fetch(trackingURL, val)
	if err != nil {
		return nil, err
	}

	results := []status{}

	doc.Find("div .tracking-invoice-block-detail li").Each(func(i int, s *goquery.Selection) {
		item := s.Find("div .item").Text()
		date := s.Find("div .date").Text()
		name := s.Find("div .name").Text()

		// 日付の書式を変更
		pt, _ := time.Parse("01月02日 15:04", date)
		date = pt.Format("01/02 15:04")

		results = append(results, status{
			date:    date,
			message: item,
			office:  name,
		})
	})

	if len(results) == 0 {
		return nil, createNotFoundError(trackingNumber)
	}

	return results, nil
}
