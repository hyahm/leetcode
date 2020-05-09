package main

import (
	"fmt"
	"strconv"
	"strings"
)

// You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example:

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) ThisValue(n int) {
	ln.Val = n
}

func (ln *ListNode) NextValue(n *ListNode) {
	ln.Next = n
	return
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := l1.reverse() + l2.reverse()
	newln := &ListNode{}
	thisNum := sum % 10
	newln.ThisValue(thisNum)
	//
	thisNum = sum / 10
	var next *ListNode
	next = newln
	for thisNum/10 > 0 {
		ln := thisNum % 10
		next.Next = &ListNode{Val: ln}
		next = next.Next
		thisNum = thisNum / 10
	}
	if thisNum > 0 {
		next.Next = &ListNode{Val: thisNum}
	}

	newln.print()

	return nil
}

func (ln *ListNode) print() {
	// 打印测试
	if ln == nil {
		return
	}
	fmt.Println(ln.Val)
	for {

		ln = ln.Next
		if ln == nil {
			break
		}
		fmt.Println(ln.Val)
	}
}

func (ln *ListNode) reverse() int {
	// 返回整数
	if ln == nil {
		return 0
	}

	l := make([]string, 0)
	l = append(l, fmt.Sprintf("%d", ln.Val))
	for {
		ln = ln.Next
		if ln == nil {
			break
		}
		l = append(l, fmt.Sprintf("%d", ln.Val))
	}

	ss := strings.Join(l, "")

	i, _ := strconv.Atoi(ss)
	return i
}

func main() {
	l1_3 := &ListNode{
		Val: 3,
	}
	l1_4 := &ListNode{
		Val: 4,
	}
	l1 := &ListNode{
		Val: 2,
	}

	l1.NextValue(l1_4)
	l1_4.NextValue(l1_3)

	l2_4 := &ListNode{
		Val: 4,
	}
	l2_6 := &ListNode{
		Val: 6,
	}
	l2 := &ListNode{
		Val: 5,
	}

	l2.NextValue(l2_6)
	l2_6.NextValue(l2_4)
	addTwoNumbers(l1, l2)
}
