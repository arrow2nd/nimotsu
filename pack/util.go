package pack

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// fetch : Webサイトを取得
func fetch(url string, val url.Values) (*goquery.Document, error) {
	res, err := http.PostForm(url, val)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("access failed (%d : %s)", res.StatusCode, res.Status)
	}

	return goquery.NewDocumentFromReader(res.Body)
}

// removeConsecutiveSpace : 連続したスペースを削除
func removeConsecutiveSpace(str string) string {
	str = strings.TrimSpace(str)
	rep := regexp.MustCompile(`[\s　]+`)
	return rep.ReplaceAllString(str, " ")
}

// createNotFoundError : 見つからなかったときのエラーを作成
func createNotFoundError(n string) error {
	return fmt.Errorf("couldn't find the package (%s)", n)
}
