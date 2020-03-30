package problem0021

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 任何一条链为空就返回另一条
	if l1 == nil {return l2}else if l2 == nil {return l1}

	var head, node *ListNode
	// 设置头结点
	if l1.Val < l2.Val {
		head = l1
		node = l1
		l1 = l1.Next
	} else {
		head = l2
		node = l2
		l2 = l2.Next
	}

	// 循环交替比较大小，直到一条链为空
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}

	// 连接剩余链
	if l1 == nil {
		node.Next = l2
	}

	if l2 == nil {
		node.Next = l1
	}

	return head
}
