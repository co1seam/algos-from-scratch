// package threesum provides a solution for the "Three Sum" problem.
package threesum

import (
	"sort"
)

// ThreeSum finds all unique triplets that sum to zero.
// the input slice is sorted in the process. The result contains no duplicates.
// it returns an empty slice if no such triplets exists.
//
// Time complexity: O(n^2)
// Space complexity: O(log n)
func ThreeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}

	sort.Ints(nums)

	n := len(nums)
	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]
		left, right := i+1, n-1

		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})

				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
