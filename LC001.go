package main

/**

1. Two Sum

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
 */

func twoSum(nums []int, target int) []int  {
	if len(nums) < 2 {
		return nil
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func twoSum2(nums []int, target int) []int  {

	if len(nums) < 2 {
		return nil
	}

	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		l := target - nums[i]
		if e, ok := m[l]; ok {
			return []int{i, e}
		} else {
			m[nums[i]] = i
		}
	}


	return nil
}

func main()  {

}


