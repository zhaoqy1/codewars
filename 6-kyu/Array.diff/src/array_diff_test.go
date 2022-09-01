package src

import (
	"reflect"
	"testing"
)

func TestArrayDiff(t *testing.T) {
	type args struct {
		a, b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"", args{[]int{1, 2}, []int{1}}, []int{2}},
		{"", args{[]int{1, 2, 2}, []int{1}}, []int{2, 2}},
		{"", args{[]int{1, 2, 2}, []int{2}}, []int{1}},
		{"", args{[]int{1, 2, 2}, []int{}}, []int{1, 2, 2}},
		{"", args{[]int{}, []int{1, 2}}, []int{}},
		{"", args{[]int{1, 2, 3}, []int{1, 2}}, []int{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayDiff(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayDiff() = %v, want %v", got, tt.want)
			}
		})
	}
}
