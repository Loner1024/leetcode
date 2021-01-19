/*
 * @lc app=leetcode.cn id=1486 lang=golang
 *
 * [1486] 数组异或操作
 */
package main

import "fmt"

// @lc code=start
func xorOperation(n int, start int) int {
	eor := 0
	for i := 0; i < n; i++ {
		eor ^= start + 2*i
	}
	return eor
}

// @lc code=end

func main() {
	fmt.Println(xorOperation(10, 5))
}
