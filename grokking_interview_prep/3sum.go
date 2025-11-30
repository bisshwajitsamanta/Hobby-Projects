package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	// Sorting arrays
	var result [][]int
	sort.Ints(nums)

	// Outer loop for i++
	for i := 0; i < len(nums)-2; i++ {
		// skip duplicates for i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1

		for j < k {
			sum := nums[i] + nums[j] + nums[k]

			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				// skip duplicates for j
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				// skip duplicates for k
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}

		}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))

}
