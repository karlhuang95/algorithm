package longestValidParentheses

func longestValidParentheses(s string) int {
	stack := []int{}
	stack = append(stack, -1)
	max := 0
	for i, c := range s {
		if c == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1] // 弹出栈顶元素
			if len(stack) == 0 {
				stack = append(stack, i)
			}
			// 当前元素的索引与栈顶元素作差，获取最近的括号匹配数
			max = Max(max, i - stack[len(stack)-1])
		}
	}
	return max
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


