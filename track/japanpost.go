package track

import (
	"fmt"
	"log"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// JapanPost 日本郵便を追跡
func JapanPost(tracknumber string) {
	val := url.Values{}
	val.Add("searchKind", "S002")
	val.Add("locale", "ja")
	val.Add("reqCodeNo1", tracknumber)

	doc, err := fetchBody("https://trackings.post.japanpost.jp/services/srv/search/direct", val)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("[summary='配達状況詳細'] td").Each(func(i int, s *goquery.Selection) {
		t := s.Text()
		fmt.Printf("%d / %s\n", i, t)
	})

	doc.Find("[summary='履歴情報'] td").Each(func(i int, s *goquery.Selection) {
		t := s.Text()
		fmt.Printf("%d / %s\n", i, t)
	})
}
