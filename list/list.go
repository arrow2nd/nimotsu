package list

import (
	"errors"

	"github.com/arrow2nd/nimotsu/pack"
)

// List : 荷物リスト
type List struct {
	dir      string
	packages []pack.Package
}

// New : 生成
func New() *List {
	return &List{
		dir:      "",
		packages: []pack.Package{},
	}
}

// Get : リストを取得
func (l *List) Get() []pack.Package {
	return l.packages
}

// Clear : リストをクリア
func (l *List) Clear() {
	l.packages = []pack.Package{}
}

// IsEmpty : リストが空かどうか
func (l *List) IsEmpty() bool {
	return len(l.packages) == 0
}

// AddItem : 荷物をリストに追加
func (l *List) AddItem(pkgs *pack.Package) {
	l.packages = append(l.packages, *pkgs)
}

// RemoveItem : 荷物をリストから削除
func (l *List) RemoveItem(number string) error {
	new := []pack.Package{}

	for _, pkg := range l.packages {
		if pkg.Number != number {
			new = append(new, pkg)
		}
	}

	if len(new) == len(l.packages) {
		return errors.New("no tracking number found")
	}

	l.packages = new
	return nil
}

// Exists : リスト内に存在するか
func (l *List) Exists(number string) bool {
	for _, pkg := range l.packages {
		if pkg.Number == number {
			return true
		}
	}

	return false
}

// ChangeComment : コメントを変更する
func (l *List) ChangeComment(number, comment string) error {
	for i := 0; i < len(l.packages); i++ {
		if l.packages[i].Number == number {
			l.packages[i].Comment = comment
			return nil
		}
	}

	return errors.New("no tracking number found")
}
