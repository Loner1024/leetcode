/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
// @lc code=start
func isValid(s string) bool {
	var stack []string
	for _, v := range s {
		switch string(v) {
		case "(", "{", "[":
			stack = append(stack, string(v))
		case ")":
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == "(" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case "}":
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == "{" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case "]":
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == "[" {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

// @lc code=end
