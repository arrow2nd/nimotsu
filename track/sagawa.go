package track

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// TrackBySagawa 佐川急便を追跡
func (t *Track) TrackBySagawa() {
	const (
		sagawaUrl = "https://k2k.sagawa-exp.co.jp/p/web/okurijosearch.do"
		fieldMax  = 3
	)

	t.company = "佐川急便"

	val := url.Values{}
	val.Add("okurijoNo", t.number)

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

			t.statuses = append(t.statuses, Status{
				Date:    fmt.Sprintf("%d/%s", time.Now().Year(), field[1]),
				Message: field[0][3:], // 先頭の文字を削除
				Office:  field[2],
			})
		})
	})
}
