/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
package main

func rotate(nums []int, k int) {
	start := nums[len(nums)-k:]
	end := nums[:len(nums)-k]
	nums = append(start, end...)
}

func main() {
	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
	rotate([]int{-1, -100, 3, 99}, 2)
}

// @lc code=end
