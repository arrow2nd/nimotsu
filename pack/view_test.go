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
	tests := []struct {
		name   string
		fields fields
		want   [][]string
	}{
		{
			name: "表データの作成",
			fields: fields{
				carrier: "わんわん運輸",
				number:  "0123456789",
				comment: "お肉",
				statuses: []status{{
					date:    "2021/11/23 00:00",
					message: "荷物受付",
					office:  "わんわん営業所",
				}},
			},
			want: [][]string{{
				"0123456789",
				"お肉",
				"わんわん運輸",
				"荷物受付",
				"2021/11/23 00:00",
				"わんわん営業所",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New(tt.fields.carrier, tt.fields.number, tt.fields.comment)
			p.statuses = tt.fields.statuses
			if got := p.CreateViewData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PackInfo.CreateViewData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPackInfo_View(t *testing.T) {
	type fields struct {
		carrier  string
		number   string
		comment  string
		statuses []status
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "データが無い時にエラーを返すか",
			fields:  fields{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New(tt.fields.carrier, tt.fields.number, tt.fields.comment)
			p.statuses = tt.fields.statuses
			if err := p.View(); (err != nil) != tt.wantErr {
				t.Errorf("PackInfo.View() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
