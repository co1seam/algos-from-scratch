// package twosum provides a solution for the "Two Sum" problem.
package twosum

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
