/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func reverseList(head *ListNode) *ListNode {
// 	var pre *ListNode
// 	cur := head
// 	for cur != nil {
// 		next := cur.Next
// 		cur.Next = pre
// 		pre = cur
// 		cur = next
// 	}
// 	return pre
// }

// 递归实现
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	rev := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return rev
}

// @lc code=end

