package list

import (
	"fmt"
)

// List : 荷物リスト
type List struct {
	dir   string
	items []Item
}

// Item : 荷物
type Item struct {
	Carrier string
	Number  string
	Comment string
}

// New : 生成
func New() *List {
	return &List{
		items: []Item{},
	}
}

// Get : リストを取得
func (l *List) Get() []Item {
	return l.items
}

// Clear : リストをクリア
func (l *List) Clear() {
	l.items = []Item{}
}

// IsEmpty : リストが空かどうか
func (l *List) IsEmpty() bool {
	return len(l.items) == 0
}

// AddItem : 荷物をリストに追加
func (l *List) AddItem(item *Item) {
	l.items = append(l.items, *item)
}

// RemoveItem : 荷物をリストから削除
func (l *List) RemoveItem(number string) error {
	new := []Item{}

	for _, item := range l.items {
		if item.Number != number {
			new = append(new, item)
		}
	}

	if len(new) == len(l.items) {
		return fmt.Errorf("no tracking number found")
	}

	l.items = new
	return nil
}

// Exists : リスト内に存在するか
func (l *List) Exists(number string) bool {
	for _, item := range l.items {
		if item.Number == number {
			return true
		}
	}

	return false
}

// ChangeComment : コメントを変更する
func (l *List) ChangeComment(number, comment string) error {
	for i := 0; i < len(l.items); i++ {
		if l.items[i].Number == number {
			l.items[i].Comment = comment
			return nil
		}
	}

	return fmt.Errorf("no tracking number found")
}
