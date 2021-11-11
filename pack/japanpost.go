package pack

import (
	"log"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// trackByJapanPost 日本郵便を追跡
func (p *PackInfo) trackByJapanPost() {
	const (
		japanpostUrl = "https://trackings.post.japanpost.jp/services/srv/search/direct"
		fieldMax     = 6
	)

	val := url.Values{}
	val.Add("searchKind", "S002")
	val.Add("locale", "ja")
	val.Add("reqCodeNo1", p.number)

	doc, err := fetchBody(japanpostUrl, val)
	if err != nil {
		log.Fatal(err)
	}

	var field [fieldMax]string

	doc.Find("[summary='履歴情報'] td").Each(func(i int, s *goquery.Selection) {
		// 配達状況を追加
		if (i+1)%fieldMax == 0 {
			p.statuses = append(p.statuses, status{
				date:    field[0],
				message: field[1],
				office:  field[4] + " " + field[3],
			})
		}

		field[i%fieldMax] = s.Text()
	})
}
