/*
 * @lc app=leetcode.cn id=240 lang=golang
 *
 * [240] 搜索二维矩阵 II
 */
package main

import "fmt"

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 {
		return false
	}
	row, col := 0, len(matrix[0])-1
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			col -= 1
		} else {
			row += 1
		}
	}
	return false
}

// @lc code=end

func main() {
	arr := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	fmt.Println(searchMatrix(arr, 5))
}
