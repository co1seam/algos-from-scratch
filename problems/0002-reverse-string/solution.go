// package revstr provides a solution for the "Reverse String" problem.
package revstr

// ReverseString reverses the order of elements in the given byte slice in place.
// it does not allocate new slice, the function modifies the underlying array of the slice.
// no changes are made, if the slice are empty or contains a single element.
//
// Time complexity: O(n)
// Space complexity: O(n)
func ReverseString(s []byte) {
	left, right := 0, len(s)-1

	for left <= right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
