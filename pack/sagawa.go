package pack

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// trackBySagawa 佐川急便を追跡
func (p *PackInfo) trackBySagawa() {
	const (
		sagawaUrl = "https://k2k.sagawa-exp.co.jp/p/web/okurijosearch.do"
		fieldMax  = 3
	)

	val := url.Values{}
	val.Add("okurijoNo", p.number)

	doc, err := fetchBody(sagawaUrl, val)
	if err != nil {
		log.Fatal(err)
	}

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

			p.statuses = append(p.statuses, status{
				date:    fmt.Sprintf("%d/%s", time.Now().Year(), field[1]),
				message: field[0][3:], // 先頭の文字を削除
				office:  field[2],
			})
		})
	})
}
