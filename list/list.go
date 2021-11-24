package list

import (
	"fmt"
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

// Get リストを取得
func (l *List) Get() []Item {
	return l.items
}

// Clear リストをクリア
func (l *List) Clear() {
	l.items = []Item{}
}

// AddItem 荷物をリストに追加
func (l *List) AddItem(item *Item) {
	l.items = append(l.items, *item)
}

// RemoveItem 荷物をリストから削除
func (l *List) RemoveItem(number string) error {
	new := []Item{}

	for _, item := range l.items {
		if item.Number != number {
			new = append(new, item)
		}
	}

	if len(new) == len(l.items) {
		return fmt.Errorf("not found")
	}

	l.items = new
	return nil
}

// Exists リスト内に存在するか
func (l *List) Exists(number string) bool {
	for _, item := range l.items {
		if item.Number == number {
			return true
		}
	}

	return false
}
