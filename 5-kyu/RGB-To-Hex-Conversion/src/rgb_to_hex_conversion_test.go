package src

import (
	"testing"
)

func TestRGB(t *testing.T) {
	type args struct {
		r, g, b int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{0, 0, 0}, "000000"},
		{"", args{1, 2, 3}, "010203"},
		{"", args{255, 255, 255}, "FFFFFF"},
		{"", args{254, 253, 252}, "FEFDFC"},
		{"", args{-20, 275, 125}, "00FF7D"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RGB(tt.args.r, tt.args.g, tt.args.b); got != tt.want {
				t.Errorf("RGB() = %v, want %v", got, tt.want)
			}
		})
	}
}
