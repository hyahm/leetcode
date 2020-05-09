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
	"reflect"
	"testing"
)

func twoSum(nums []int, target int) []int {

	for i, v := range nums {

		for j, s := range nums[i+1:] {
			if v+s == target {
				return []int{i, i + j + 1}
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
	// 3, 2, 4
	for i, v := range nums {

		if j, ok := m[target-v]; ok && i != j {
			fmt.Println("now", target-v)
			fmt.Println(i, j)
			fmt.Println([]int{i, j})
			return []int{i, j}
		}
	}
	return nil
}

type data struct {
	nums   []int
	target int
	expect []int
	title  string
}

var n = []data{
	{
		nums:   []int{2, 17, 11, 15, 6, 3},
		target: 9,
		expect: []int{4, 5},
		title:  "test1",
	},
	{
		nums:   []int{-1, -2, -3, -4, -5},
		target: -8,
		expect: []int{2, 4},
		title:  "test2",
	},
	{
		nums:   []int{3, 2, 4},
		target: 6,
		expect: []int{1, 2},
		title:  "test3",
	},
}

func TestTwoSumHash(t *testing.T) {

	for _, v := range n {
		if !reflect.DeepEqual(twoSumHash(v.nums, v.target), v.expect) {
			t.Errorf("not equel by %s", v.title)
		}
	}
	// fmt.Println(twoSum(nums, target))

}

func TestTwoSum(t *testing.T) {

	for _, v := range n {
		if !reflect.DeepEqual(twoSum(v.nums, v.target), v.expect) {
			t.Errorf("not equel by %s", v.title)
		}
	}
}
