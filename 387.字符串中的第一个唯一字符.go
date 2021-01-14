/*
 * @lc app=leetcode.cn id=387 lang=golang
 *
 * [387] 字符串中的第一个唯一字符
 */

// @lc code=start
package main

import "fmt"

func firstUniqChar(s string) int {
	for i, s1 := range s {
		sign := 0
		for j, s2 := range s {
			if j != i && s1 == s2 {
				sign = 1
				break
			}
		}
		if sign == 0 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
	fmt.Println(firstUniqChar("loveleetcode"))
}

// @lc code=end
