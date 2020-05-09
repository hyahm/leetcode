package main

import "fmt"

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example:

// Given nums = [2, 7, 11, 15], target = 9,

// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

func twoSumHash(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		m[v] = k
	}
	for i, v := range nums {
		if j, ok := m[target-v]; ok && i != j {
			return []int{i, j}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 17, 11, 15, 6, 3}
	target := 9
	fmt.Println(twoSumHash(nums, target))
}
