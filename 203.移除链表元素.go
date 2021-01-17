/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 */
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 循环实现
// func removeElements(head *ListNode, val int) *ListNode {
// 	if head == nil {
// 		return head
// 	}
// 	for head.Val == val && head != nil {
// 		head = head.Next
// 		if head == nil {
// 			return head
// 		}
// 	}
// 	pre := head
// 	for pre.Next != nil {
// 		if pre.Next.Val == val {
// 			delNode := pre.Next
// 			pre.Next = delNode.Next
// 			delNode.Next = nil
// 		} else {
// 			pre = pre.Next
// 		}
// 	}
// 	return head
// }

// 递归实现
func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return removeElements(head.Next, val)
	} else {
		head.Next = removeElements(head.Next, val)
		return head
	}
}

// @lc code=end
