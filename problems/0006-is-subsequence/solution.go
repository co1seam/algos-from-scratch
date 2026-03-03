// package subseq provides a solution for the "Is Subsequence" problem.
package subseq

// IsSubsequence checks wheter string 's' is a subsequence of string 't'.
// it return true, if 's' is empty  or if every characters of 's' appers in 't' in order
// otherwise it returns false.
//
// Time complexity: O(n)
// Time complextity: O(1)
func IsSubsequence(s, t string) bool {
	if s == "" {
		return true
	}

	bottom := 0

	for top := 0; top < len(t); top++ {
		if s[bottom] == t[top] {
			bottom++
		}
		if bottom > len(s)-1 {
			return true
		}
	}

	return false
}
