package pack

import (
	"net/url"
	"reflect"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func Test_fetchBody(t *testing.T) {
	type args struct {
		url string
		val url.Values
	}
	tests := []struct {
		name    string
		args    args
		want    *goquery.Document
		wantErr bool
	}{
		{
			name: "正しくエラーが返るか",
			args: args{
				url: "https://hoge.test",
				val: map[string][]string{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetchBody(tt.args.url, tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fetchBody() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeConsecutiveSpace(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "前後の半角スペースを削除",
			args: args{
				str: "  あいうえお  ",
			},
			want: "あいうえお",
		},
		{
			name: "前後の全角スペースを削除",
			args: args{
				str: "　　あいうえお　　　",
			},
			want: "あいうえお",
		},
		{
			name: "連続した半角スペースを1つの半角スペースに置換",
			args: args{
				str: "あ  い  う  え  お",
			},
			want: "あ い う え お",
		},
		{
			name: "連続した全角スペースを1つの半角スペースに置換",
			args: args{
				str: "あ　　い　　う　　え　　お",
			},
			want: "あ い う え お",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeConsecutiveSpace(tt.args.str); got != tt.want {
				t.Errorf("removeConsecutiveSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
