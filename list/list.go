package list

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const filename = ".nimotsu"

// List 荷物リスト
type List struct {
	items []Item
}

// Item 荷物
type Item struct {
	Carrier string
	Number  string
	Comment string
}

// New 生成
func New() *List {
	return &List{
		items: []Item{},
	}
}

// AddItem 荷物をリストに追加
func (l *List) AddItem(item Item) {
	l.items = append(l.items, item)
}

// RemoveItem 荷物をリストから削除
func (l *List) RemoveItem(number string) {
	new := []Item{}

	for _, item := range l.items {
		if item.Number != number {
			new = append(new, item)
		}
	}

	l.items = new
}

// Save ファイルへ保存
func (l *List) Save() error {
	buf, err := yaml.Marshal(l.items)
	if err != nil {
		return err
	}

	path := getSaveFilePath()
	err = ioutil.WriteFile(path, buf, os.ModePerm)
	if err != nil {
		fmt.Println("Error: Failed to write file")
		panic(err)
	}

	return nil
}

// Load ファイルから読込
func (l *List) Load() error {
	path := getSaveFilePath()
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, &l.items)
	if err != nil {
		return err
	}

	return nil
}

func getSaveFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return filepath.Join(homeDir, filename)
}
