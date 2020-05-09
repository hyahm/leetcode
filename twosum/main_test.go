package main

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.

// You may assume that each input would have exactly one solution, and you may not use the same element twice.

// Example:

// Given nums = [2, 7, 11, 15], target = 9,

// Because nums[0] + nums[1] = 2 + 7 = 9,
// return [0, 1].

// run:
// go test -v -run=TestTwoSum leetcode/twosum
// go test -v -run=TestTwoSumHash leetcode/twosum

import (
	"fmt"
	"testing"
)

func twoSum(nums []int, target int) []int {

	for i, v := range nums {
		// 如果大于target， 直接跳过
		if v > target {
			continue
		} else if i < len(nums)-1 {
			for j, s := range nums[i+1:] {
				if v+s == target {
					return []int{i, i + j + 1}
				}
			}
		}
	}
	return nil
}

func twoSumHash(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		m[v] = k
	}
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{i, j}
		}
	}
	return nil
}

func TestTwoSumHash(t *testing.T) {
	nums := []int{2, 17, 11, 15, 6, 3}
	target := 9
	// fmt.Println(twoSum(nums, target))

	fmt.Println(twoSumHash(nums, target))
}

func TestTwoSum(t *testing.T) {
	nums := []int{2, 17, 11, 15, 6, 3}
	target := 9
	// fmt.Println(twoSum(nums, target))

	fmt.Println(twoSum(nums, target))
}
