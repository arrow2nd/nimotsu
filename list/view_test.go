package list

import (
	"reflect"
	"testing"
)

func TestList_createTableData(t *testing.T) {
	type fields struct {
		items []Item
	}
	tests := []struct {
		name   string
		fields fields
		want   [][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				items: tt.fields.items,
			}
			if got := l.createTableData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List.createTableData() = %v, want %v", got, tt.want)
			}
		})
	}
}
