package pack

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func Test_fetch(t *testing.T) {
	tests := []struct {
		name       string
		statusCode int
		wantErr    bool
	}{
		{
			name:       "取得できるか",
			statusCode: http.StatusOK,
			wantErr:    false,
		},
		{
			name:       "エラーが返るか（404）",
			statusCode: http.StatusNotFound,
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// テストサーバーを作成
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.statusCode)
				if tt.statusCode == http.StatusOK {
					w.Write([]byte("<html><body>Test</body></html>"))
				}
			}))
			defer ts.Close()

			got, err := fetch(ts.URL, url.Values{})
			if (err != nil) != tt.wantErr {
				t.Errorf("fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && got == nil {
				t.Errorf("no value for fetch")
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
