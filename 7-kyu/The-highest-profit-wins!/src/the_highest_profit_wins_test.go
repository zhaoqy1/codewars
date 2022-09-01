package src

import "testing"

func TestMinMax(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want [2]int
	}{
		{"", args{[]int{1, 2, 3, 4, 5}}, [2]int{1, 5}},
		{"", args{[]int{2334454, 5}}, [2]int{5, 2334454}},
		{"", args{[]int{5}}, [2]int{5, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinMax(tt.args.arr); got != tt.want {
				t.Errorf("MinMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
