package src

import (
	"testing"
)

func TestZeros(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{0}, 0},
		{"", args{6}, 1},
		{"", args{30}, 7},
		{"", args{100}, 24},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Zeros(tt.args.n); got != tt.want {
				t.Errorf("Zeros() = %v, want %v", got, tt.want)
			}
		})
	}
}
