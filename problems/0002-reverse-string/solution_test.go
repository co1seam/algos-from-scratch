package revstr

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name  string
		input []byte
		want  []byte
	}{
		{"basic case", []byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{"palindrome", []byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReverseString(tt.input)
			if !reflect.DeepEqual(tt.input, tt.want) {
				t.Errorf("ReverseString() = %v, want %v", tt.input, tt.want)
			}
		})
	}
}
