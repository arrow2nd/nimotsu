package pack

import (
	"reflect"
	"testing"
)

func TestPackInfo_CreateViewData(t *testing.T) {
	type fields struct {
		carrier  string
		number   string
		comment  string
		statuses []status
	}

	testFields := fields{
		carrier: "わんわん運輸",
		number:  "0123456789",
		comment: "お肉",
		statuses: []status{{
			date:    "2021/11/23 00:00",
			message: "荷物受付",
			office:  "わんわん営業所",
		}},
	}
	want := [][]string{{
		"0123456789",
		"お肉",
		"わんわん運輸",
		"荷物受付",
		"2021/11/23 00:00",
		"わんわん営業所",
	}}

	t.Run("表データ作成", func(t *testing.T) {
		p := New(testFields.carrier, testFields.number, testFields.comment)
		p.statuses = testFields.statuses
		if got := p.CreateViewData(); !reflect.DeepEqual(got, want) {
			t.Errorf("PackInfo.CreateViewData() = %v, want %v", got, want)
		}
	})
}

func TestPackInfo_View(t *testing.T) {
	t.Run("データが無い時にエラーを返すか", func(t *testing.T) {
		p := &Package{}
		if err := p.View(); err == nil {
			t.Errorf("PackInfo.View() error = %v, wantErr %v", err, true)
		}
	})
}
