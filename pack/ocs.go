package pack

import (
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// CarrierOCS : OCS（ANAの国際輸送サービス）
const CarrierOCS CarrierName = "OCS"

func init() {
	carriers[CarrierOCS] = &carrier{
		CarrierInfo: &CarrierInfo{
			Key:    "ocs",
			Alias:  "o",
			NameEn: "OCS Express",
		},
		tracking: trackingByOCS,
	}
}

func trackingByOCS(trackingNumber string) ([]status, error) {
	const trackingURL = "https://webcsw.ocs.co.jp/csw/ECSWG0201R00003P.do"

	val := url.Values{}
	val.Add("cwbno", trackingNumber)

	doc, err := fetch(trackingURL, val)
	if err != nil {
		return nil, err
	}

	results := []status{}
	table := doc.Find("div #chartarea table").Eq(4)

	table.Find("#chart tr").Each(func(i int, s *goquery.Selection) {
		td := s.Find("td").Map(func(i int, s *goquery.Selection) string {
			// なぜか<input>が使われているセルがあるので、属性値を見る
			if v, ok := s.Children().Attr("value"); ok {
				return v
			}
			return removeConsecutiveSpace(s.Text())
		})

		pt, _ := time.Parse("Mon 02Jan2006 15:04", td[1])
		date := pt.Format("2006/01/02 15:04")

		message := td[0]
		if td[3] != "" {
			message += " / " + td[3]
		}

		r := status{
			date:    date,
			message: message,
			office:  td[2],
		}

		results = append([]status{r}, results...)
	})

	if len(results) == 0 {
		return nil, createNotFoundError(trackingNumber)
	}

	return results, nil
}
