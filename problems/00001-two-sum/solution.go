// package twosum provides a solution for the "Two Sum" problem.
package twosum

// TwoSum finds two numbers in nums array that add up to target and return their indicies.
// it return nil if no such pair exists.
//
// Time complexity: O(n)
// Space complexity: O(n)
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		if j, ok := seen[target-num]; ok {
			return []int{j, i}
		}
		seen[num] = i
	}

	return nil
}
