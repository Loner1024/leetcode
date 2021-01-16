/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start
// 递归实现
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

// @lc code=end

