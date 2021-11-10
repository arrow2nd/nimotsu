package track

import (
	"log"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// TrackByJapanPost 日本郵便を追跡
func (t *Track) TrackByJapanPost() {
	const (
		japanpostUrl = "https://trackings.post.japanpost.jp/services/srv/search/direct"
		fieldMax     = 6
	)

	val := url.Values{}
	val.Add("searchKind", "S002")
	val.Add("locale", "ja")
	val.Add("reqCodeNo1", t.number)

	doc, err := fetchBody(japanpostUrl, val)
	if err != nil {
		log.Fatal(err)
	}

	var field [fieldMax]string

	doc.Find("[summary='履歴情報'] td").Each(func(i int, s *goquery.Selection) {
		// 履歴情報を追加
		if (i+1)%fieldMax == 0 {
			t.statuses = append(t.statuses, Status{
				Date:    field[0],
				Message: field[1],
				Memo:    field[2],
				Office:  field[4] + " " + field[3],
			})
		}

		field[i%fieldMax] = s.Text()
	})
}
