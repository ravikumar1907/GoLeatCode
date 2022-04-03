package main

/*
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedmerge(a, b *ListNode) *ListNode {
	var result *ListNode
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Val <= b.Val {
		result = a
		result.Next = sortedmerge(a.Next, b)
	} else {
		result = b
		result.Next = sortedmerge(a, b.Next)
	}
	return result
}
func mergeKLists(lists []*ListNode) *ListNode {
	var head *ListNode
	n := len(lists)
	for i := 0; i < n; i++ {
		head = sortedmerge(head, lists[i])
	}

	return head
}
