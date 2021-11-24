package list

import (
	"reflect"
	"testing"
)

func TestList_createTableData(t *testing.T) {
	items := []Item{{
		Carrier: "A",
		Number:  "123456",
		Comment: "testA",
	}}
	want := [][]string{{
		"123456",
		"testA",
		"A",
	}}
	t.Run("表データ作成", func(t *testing.T) {
		l := &List{items: items}
		if got := l.createTableData(); !reflect.DeepEqual(got, want) {
			t.Errorf("List.createTableData() = %v, want %v", got, want)
		}
	})
}
