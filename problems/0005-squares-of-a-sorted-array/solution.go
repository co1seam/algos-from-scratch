// package sortsqr provides a solution for the "Squares of a Sorted Array" problem.
package sortsqr

// SortedSqueares takes a slice of integers sorted in non non-decreasing order
// and returns a new slice  containing the squares of each num, also sorted in non-decreasing order.
//
// Time complexity: O(n)
// Space complexity: O(n)
func SortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	left, right := 0, n-1

	for i := n - 1; i >= 0; i-- {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			res[i] = nums[left] * nums[left]
			left++
		} else {
			res[i] = nums[right] * nums[right]
			right--
		}
	}

	return res
}
