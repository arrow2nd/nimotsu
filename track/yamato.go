package track

import (
	"log"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// TrackByYamato ヤマト運輸を追跡
func (t *Track) TrackByYamato() {
	const yamatoUrl = "https://toi.kuronekoyamato.co.jp/cgi-bin/tneko"

	val := url.Values{}
	val.Add("number00", "1") // 取得件数？
	val.Add("number01", t.number)

	doc, err := fetchBody(yamatoUrl, val)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div .tracking-invoice-block-detail li").Each(func(i int, s *goquery.Selection) {
		item := s.Find("div .item").Text()
		date := s.Find("div .date").Text()
		name := s.Find("div .name").Text()

		t.statuses = append(t.statuses, Status{
			Date:    date,
			Message: item,
			Office:  name,
		})
	})
}
