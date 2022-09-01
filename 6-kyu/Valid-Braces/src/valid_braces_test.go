package src

import "testing"

func TestValidBraces(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{"(){}[]"}, true},
		{"", args{"([{}])"}, true},
		{"", args{"(}"}, false},
		{"", args{"[(])"}, false},
		{"", args{"[({})](]"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidBraces(tt.args.str); got != tt.want {
				t.Errorf("ValidBraces() = %v, want %v", got, tt.want)
			}
		})
	}
}
