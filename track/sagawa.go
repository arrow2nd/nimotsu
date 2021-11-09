package track

import (
	"fmt"
	"log"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// Sagawa 佐川急便を追跡
func Sagawa(tracknumber string) {
	val := url.Values{}
	val.Add("okurijoNo", tracknumber)

	doc, err := fetchBody("https://k2k.sagawa-exp.co.jp/p/web/okurijosearch.do", val)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("table.table_basic.table_okurijo_detail2").Each(func(i int, s *goquery.Selection) {
		s.Find("tr").Each(func(i int, s *goquery.Selection) {
			title := removeConsecutiveSpace(s.Find("th").Text())
			data := removeConsecutiveSpace(s.Find("td").Text())
			fmt.Printf("%d : %s / %s\n", i, title, data)
		})

		fmt.Printf("--- %d ---\n", i)
	})
}
