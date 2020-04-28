package xhas_test

import (
	"fmt"
	"testing"

	"github.com/gmidorii/xhas"
)

func ExampleString() {
	ok := xhas.String("hoge", []string{"fuga", "hoge", "null"})
	fmt.Println(ok)
	// Output: true
}

func TestString(t *testing.T) {
	type args struct {
		dst string
		src []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "src has dst",
			args: args{
				dst: "a",
				src: []string{"a", "b", "c"},
			},
			want: true,
		},
		{
			name: "src donot has dst",
			args: args{
				dst: "d",
				src: []string{"a", "b", "c"},
			},
			want: false,
		},
		{
			name: "japanese",
			args: args{
				dst: "日",
				src: []string{"あ", "ア", "日"},
			},
			want: true,
		},
		{
			name: "multiple string",
			args: args{
				dst: "abc def ghi",
				src: []string{"xxx yyyy", "abc def ghi", "iiii zzzzz"},
			},
			want: true,
		},
		{
			name: "emoji",
			args: args{
				dst: "👍",
				src: []string{"🙇", "👍", "🏠"},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xhas.String(tt.args.dst, tt.args.src); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
