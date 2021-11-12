package pack

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

const (
	JapanPost  = "日本郵便"
	jpUrl      = "https://trackings.post.japanpost.jp/services/srv/search/direct"
	jpFieldMax = 6
)

// trackByJapanPost 日本郵便を追跡
func (p *PackInfo) trackByJapanPost() error {
	val := url.Values{}
	val.Add("searchKind", "S002")
	val.Add("locale", "ja")
	val.Add("reqCodeNo1", p.number)

	doc, err := fetchBody(jpUrl, val)
	if err != nil {
		return err
	}

	var field [jpFieldMax]string
	
	doc.Find("[summary='履歴情報'] td").Each(func(i int, s *goquery.Selection) {
		// 配達状況を追加
		if (i+1)%jpFieldMax == 0 {
			p.statuses = append(p.statuses, status{
				date:    field[0],
				message: field[1],
				office:  field[4] + " " + field[3],
			})
		}

		field[i%jpFieldMax] = s.Text()
	})

	return nil
}
