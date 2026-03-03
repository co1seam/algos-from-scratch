package subseq

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		seq    string
		subseq string
		want   bool
	}{
		{"basic case", "ahbgdc", "abc", true},
		{"is not subsequence", "ahbgdc", "axc", false},
		{"empty subsequence", "ahbhdc", "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsSubsequence(tt.subseq, tt.seq); got != tt.want {
				t.Errorf("IsSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
