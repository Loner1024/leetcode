/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
	for i, v := range nums {
		for j, t := range nums {
			if v == t && i != j {
				return true
			}
		}
	}
	return false
}

// @lc code=end

