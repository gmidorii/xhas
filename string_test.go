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

func TestStringNotOne(t *testing.T) {
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
			name: "src donot has dst",
			args: args{
				dst: "abc",
				src: []string{"def", "hoge", "fuga"},
			},
			want: true,
		},
		{
			name: "src has dst",
			args: args{
				dst: "abc",
				src: []string{"def", "hoge", "abc"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xhas.StringNotOne(tt.args.dst, tt.args.src); got != tt.want {
				t.Errorf("StringNotOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringAll(t *testing.T) {
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
			name: "src match all dst",
			args: args{
				dst: "abc",
				src: []string{"abc", "abc", "abc"},
			},
			want: true,
		},
		{
			name: "src match one dst",
			args: args{
				dst: "abc",
				src: []string{"def", "abc", "hoge"},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xhas.StringAll(tt.args.dst, tt.args.src); got != tt.want {
				t.Errorf("StringAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringPre(t *testing.T) {
	type args struct {
		prefix string
		src    []string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "src has match prefix",
			args: args{
				prefix: "yyy",
				src:    []string{"xxx_a", "yyy_b", "zzz_c"},
			},
			want:  "yyy_b",
			want1: true,
		},
		{
			name: "return first match src",
			args: args{
				prefix: "yyy",
				src:    []string{"xxx_a", "yyy_b", "yyy_c", "yyy_d"},
			},
			want:  "yyy_b",
			want1: true,
		},
		{
			name: "not found",
			args: args{
				prefix: "aaa",
				src:    []string{"xxx_a", "yyy_b", "zzz_c"},
			},
			want:  "",
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := xhas.StringPre(tt.args.prefix, tt.args.src)
			if got != tt.want {
				t.Errorf("StringPre() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("StringPre() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStringPartRune(t *testing.T) {
	type args struct {
		dst rune
		src []string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "src has dst",
			args: args{
				dst: 'a',
				src: []string{"aaa", "bbb", "ccc"},
			},
			want:  "aaa",
			want1: true,
		},
		{
			name: "first match",
			args: args{
				dst: 'a',
				src: []string{"aaa1", "aaa2", "aaa3"},
			},
			want:  "aaa1",
			want1: true,
		},
		{
			name: "japanese",
			args: args{
				dst: '日',
				src: []string{"あいうえお", "日本", "こんにちは"},
			},
			want:  "日本",
			want1: true,
		},
		{
			name: "emoji",
			args: args{
				dst: '👍',
				src: []string{"🙇🙇🙇🙇🙇🙇", "🏡👍🏡🏡", "🏠"},
			},
			want:  "🏡👍🏡🏡",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := xhas.StringPartRune(tt.args.dst, tt.args.src)
			if got != tt.want {
				t.Errorf("StringPartRune() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("StringPartRune() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
