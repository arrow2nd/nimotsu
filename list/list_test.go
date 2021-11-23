package list

import (
	"reflect"
	"testing"
)

func TestList_Get(t *testing.T) {
	t.Run("取得できるか", func(t *testing.T) {
		l := New()
		want := []Item{
			{
				Carrier: "A",
				Number:  "0123456789",
				Comment: "testA",
			},
			{
				Carrier: "B",
				Number:  "9876543210",
				Comment: "testB",
			},
		}
		l.items = want

		if got := l.Get(); !reflect.DeepEqual(got, want) {
			t.Errorf("List.Get() = %v, want %v", got, want)
		}
	})
}

func TestList_Clear(t *testing.T) {
	t.Run("全て削除できるか", func(t *testing.T) {
		l := New()
		l.items = []Item{{
			Carrier: "test",
			Number:  "0123456789",
			Comment: "test!",
		}}
		l.Clear()

		want := []Item{}

		if got := l.Get(); !reflect.DeepEqual(got, want) {
			t.Errorf("List.Get() = %v, want %v", got, want)
		}
	})
}

func TestList_AddItem(t *testing.T) {
	t.Run("追加できるか", func(t *testing.T) {
		l := New()
		want := []Item{
			{
				Carrier: "C",
				Number:  "12345",
				Comment: "testC",
			},
			{
				Carrier: "D",
				Number:  "67890",
				Comment: "testD",
			},
		}
		l.items = want[:1]
		l.AddItem(&want[1])

		if got := l.Get(); !reflect.DeepEqual(got, want) {
			t.Errorf("List.Get() = %v, want %v", got, want)
		}
	})
}

func TestList_RemoveItem(t *testing.T) {
	type fields struct {
		items []Item
	}
	type args struct {
		number string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "削除できるか",
			fields: fields{
				items: []Item{
					{
						Carrier: "A",
						Number:  "012345",
						Comment: "testA",
					},
					{
						Carrier: "B",
						Number:  "678901",
						Comment: "testB",
					},
				},
			},
			args: args{
				number: "012345",
			},
			wantErr: false,
		},
		{
			name: "指定された番号が存在しない場合エラーが返るか",
			fields: fields{
				items: []Item{{
					Carrier: "B",
					Number:  "678901",
					Comment: "testB",
				},
				},
			},
			args: args{
				number: "123456",
			},
			wantErr: true,
		},
		{
			name: "データが存在しない場合エラーが返るか",
			fields: fields{
				items: []Item{},
			},
			args: args{
				number: "123456",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := New()
			l.items = tt.fields.items
			if err := l.RemoveItem(tt.args.number); (err != nil) != tt.wantErr {
				t.Errorf("List.RemoveItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestList_Exists(t *testing.T) {
	type fields struct {
		items []Item
	}
	type args struct {
		number string
	}
	testItems := []Item{
		{
			Carrier: "A",
			Number:  "123456",
			Comment: "testA",
		},
		{
			Carrier: "B",
			Number:  "789012",
			Comment: "testB",
		},
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "存在する番号",
			fields: fields{
				items: testItems,
			},
			args: args{
				number: "789012",
			},
			want: true,
		},
		{
			name: "存在しない番号",
			fields: fields{
				items: testItems,
			},
			args: args{
				number: "000000",
			},
			want: false,
		},
		{
			name: "データが存在しない",
			fields: fields{
				items: []Item{},
			},
			args: args{
				number: "123456",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &List{
				items: tt.fields.items,
			}
			if got := l.Exists(tt.args.number); got != tt.want {
				t.Errorf("List.Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}
