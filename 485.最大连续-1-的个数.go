/*
 * @lc app=leetcode.cn id=485 lang=golang
 *
 * [485] 最大连续 1 的个数
 */
package main

import "fmt"

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	cur := 0
	for _, v := range nums {
		if v == 1 {
			cur += 1
		} else {
			cur = 0
		}
		if cur > max {
			max = cur
		}
	}
	return max
}

// @lc code=end

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1}))
}
