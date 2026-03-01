package sortsqr

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"basic case", []int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{"duplicates", []int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortedSquares(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
