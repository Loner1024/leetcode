/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			cur.Next = l2
			cur = cur.Next
			l2 = l2.Next
		} else {
			cur.Next = l1
			cur = cur.Next
			l1 = l1.Next
		}
	}
	for l1 != nil {
		cur.Next = l1
		l1 = nil
		// cur = cur.Next
		// l1 = l1.Next
	}
	for l2 != nil {
		cur.Next = l2
		l2 = nil
		// cur = cur.Next
		// l2 = l2.Next
	}
	return dummy.Next
}

// @lc code=end

