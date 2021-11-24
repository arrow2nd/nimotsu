package pack

import (
	"fmt"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	// SagawaExpress 配送業者名
	SagawaExpress = "佐川急便"
	sgUrl         = "https://k2k.sagawa-exp.co.jp/p/web/okurijosearch.do"
	sgFieldMax    = 3
)

// trackBySagawa 佐川急便を追跡
func (p *PackInfo) trackBySagawa() error {
	val := url.Values{}
	val.Add("okurijoNo", p.number)

	doc, err := fetchBody(sgUrl, val)
	if err != nil {
		return err
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

			var field [sgFieldMax]string
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
		return fmt.Errorf("couldn't find the package (" + p.number + ")")
	}

	p.statuses = results
	return nil
}
