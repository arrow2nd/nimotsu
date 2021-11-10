package track

import (
	"fmt"
	"log"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// TrackByYamato ヤマト運輸を追跡
func (t *Track) TrackByYamato() {
	val := url.Values{}
	val.Add("number00", "1") // 取得件数？
	val.Add("number01", t.number)

	doc, err := fetchBody("https://toi.kuronekoyamato.co.jp/cgi-bin/tneko", val)
	if err != nil {
		log.Fatal(err)
	}

	// 荷物情報を抽出
	doc.Find("div .tracking-invoice-block-summary li").Each(func(i int, s *goquery.Selection) {
		item := s.Find("div .item").Text()
		data := s.Find("div .data").Text()
		fmt.Printf("%d : %s / %s \n", i, item, data)
	})

	// 配達状況を抽出
	doc.Find("div .tracking-invoice-block-detail li").Each(func(i int, s *goquery.Selection) {
		item := s.Find("div .item").Text()
		date := s.Find("div .date").Text()
		name := s.Find("div .name").Text()
		fmt.Printf("%d : %s / %s / %s\n", i, item, date, name)
	})
}
