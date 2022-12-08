package pack

import (
	"net/url"
	"testing"
)

func Test_fetchBody(t *testing.T) {
	type args struct {
		url string
		val url.Values
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "取得できるか",
			args: args{
				url: "http://example.com/",
				val: url.Values{},
			},
			wantErr: false,
		},
		{
			name: "エラーが返るか（404）",
			args: args{
				url: "https://httpstat.us/404",
				val: url.Values{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetch(tt.args.url, tt.args.val)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got == nil {
				t.Errorf("no value for fetchBody")
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
