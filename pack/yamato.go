package pack

import (
	"fmt"
	"net/url"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	// YamatoTransport 配送業者名
	YamatoTransport = "ヤマト運輸"
	ymUrl           = "https://toi.kuronekoyamato.co.jp/cgi-bin/tneko"
)

// trackByYamato ヤマト運輸を追跡
func (p *PackInfo) trackByYamato() error {
	val := url.Values{}
	val.Add("number00", "1") // 取得件数？
	val.Add("number01", p.number)

	doc, err := fetchBody(ymUrl, val)
	if err != nil {
		return err
	}

	var results []status

	doc.Find("div .tracking-invoice-block-detail li").Each(func(i int, s *goquery.Selection) {
		item := s.Find("div .item").Text()
		date := s.Find("div .date").Text()
		name := s.Find("div .name").Text()

		// 日付の書式を変更
		pt, _ := time.Parse("01月02日 15:04", date)
		date = fmt.Sprintf("%d/%s", time.Now().Year(), pt.Format("01/02 15:04"))

		results = append(results, status{
			date:    date,
			message: item,
			office:  name,
		})
	})

	if len(results) == 0 {
		return fmt.Errorf("couldn't find the package (" + p.number + ")")
	}

	p.statuses = results
	return nil
}
