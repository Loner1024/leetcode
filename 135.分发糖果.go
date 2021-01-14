/*
 * @lc app=leetcode.cn id=135 lang=golang
 *
 * [135] 分发糖果
 */

// @lc code=start
package main

import "fmt"

func candy(ratings []int) int {
	sum := len(ratings)
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {

		}
	}
}

func main() {
	fmt.Println(candy([]int{1, 3, 2, 2, 1}))
	fmt.Println(candy([]int{1, 2, 2}))
}

// @lc code=end
