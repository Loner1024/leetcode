/*
 * @lc app=leetcode.cn id=54 lang=golang
 *
 * [54] 螺旋矩阵
 */
// @lc code=start
func spiralOrder(matrix [][]int) []int {
	top, bootom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1
	result := []int{}
	for {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++
		if top > bootom {
			break
		}
		for i := top; i <= bootom; i++ {
			result = append(result, matrix[i][right])
		}
		right--
		if right < left {
			break
		}
		for i := right; i >= left; i-- {
			result = append(result, matrix[bootom][i])
		}
		bootom--
		if top > bootom {
			break
		}
		for i := bootom; i >= top; i-- {
			result = append(result, matrix[i][left])
		}
		left++
		if right < left {
			break
		}
	}
	return result
}

// @lc code=end
