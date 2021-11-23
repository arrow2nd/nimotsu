package pack

import (
	"testing"
)

func Test_fetchBody(t *testing.T) {
	t.Run("正しくエラーが返るか", func(t *testing.T) {
		_, err := fetchBody("https://hoge.test", map[string][]string{})
		if err == nil {
			t.Errorf("fetchBody() error = %v, wantErr %v", err, true)
			return
		}
	})
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
