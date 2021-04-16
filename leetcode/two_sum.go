/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

package main

func twoSum(nums []int, target int) []int {
	var ret []int

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			// fmt.Println("i: ", nums[i], "j: ", nums[j])
			// Additional special case -> the same number cannot be used
			if (nums[i]+nums[j]) == target && (i != j) {
				ret = append(ret, i)
				ret = append(ret, j)
				return ret
			}
		}
	}

	return ret

}
