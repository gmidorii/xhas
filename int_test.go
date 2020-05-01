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

func TestIntOdd(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "src has odd",
			args: args{
				src: []int{2, 4, 1, 4},
			},
			want: true,
		},
		{
			name: "src donot has odd",
			args: args{
				src: []int{2, 4, 10, 4},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xhas.IntOdd(tt.args.src); got != tt.want {
				t.Errorf("IntOdd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntEven(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "src has even",
			args: args{
				src: []int{2, 4, 1, 4},
			},
			want: true,
		},
		{
			name: "src donot has even",
			args: args{
				src: []int{1, 3, 5, 7},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xhas.IntEven(tt.args.src); got != tt.want {
				t.Errorf("IntEven() = %v, want %v", got, tt.want)
			}
		})
	}
}
