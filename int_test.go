package xhas_test

import (
	"testing"

	"github.com/gmidorii/xhas"
)

func TestInt(t *testing.T) {
	type args struct {
		dst int
		src []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "src has dst",
			args: args{
				dst: 1,
				src: []int{2, 3, 4, 5, 1, 7},
			},
			want: true,
		},
		{
			name: "not found",
			args: args{
				dst: 100,
				src: []int{2, 3, 4, 5, 1, 7},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xhas.Int(tt.args.dst, tt.args.src); got != tt.want {
				t.Errorf("Int() = %v, want %v", got, tt.want)
			}
		})
	}
}
