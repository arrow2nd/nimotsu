package pack

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// trackByYamato ヤマト運輸を追跡
func (p *PackInfo) trackByYamato() {
	const yamatoUrl = "https://toi.kuronekoyamato.co.jp/cgi-bin/tneko"

	val := url.Values{}
	val.Add("number00", "1") // 取得件数？
	val.Add("number01", p.number)

	doc, err := fetchBody(yamatoUrl, val)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div .tracking-invoice-block-detail li").Each(func(i int, s *goquery.Selection) {
		item := s.Find("div .item").Text()
		date := s.Find("div .date").Text()
		name := s.Find("div .name").Text()

		// 日付の書式を変更
		pt, _ := time.Parse("01月02日 15:04", date)
		date = fmt.Sprintf("%d/%s", time.Now().Year(), pt.Format("01/02 15:04"))

		p.statuses = append(p.statuses, status{
			date:    date,
			message: item,
			office:  name,
		})
	})
}
