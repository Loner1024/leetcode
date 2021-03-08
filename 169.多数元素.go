/*
 * @lc app=leetcode.cn id=169 lang=golang
 *
 * [169] 多数元素
 */
package main

import "fmt"

// @lc code=start
func majorityElement(nums []int) int {
	cur, count := nums[0], 0
	for _, v := range nums {
		if count == 0 {
			cur = v
		}
		if cur == v {
			count++
		} else {
			count--
		}
	}
	return cur
}

// @lc code=end
func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
}
