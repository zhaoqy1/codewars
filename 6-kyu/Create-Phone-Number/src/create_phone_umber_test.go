package src

import "testing"

func TestCreatePhoneNumber(t *testing.T) {
	type args struct {
		numbers [10]uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{[10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}}, "(123) 456-7890"},
		{"", args{[10]uint{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}}, "(098) 765-4321"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreatePhoneNumber(tt.args.numbers); got != tt.want {
				t.Errorf("Multiple3And5() = %v, want %v", got, tt.want)
			}
		})
	}
}
