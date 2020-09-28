package removeOuterParentheses

import "fmt"

func removeOuterParentheses(S string) string {
	res := ""
	num := 0
	index := 0
	for k, v := range S {
		if string(v) == "(" {
			num++
		} else {
			num--
		}
		if num == 1 && string(v) == "(" {
			index = k
		}
		if num == 0 {
			res += S[index+1 : k]
		}
	}
	return res

}

func removeOuterParentheses1(S string) string {
	res := ""
	num := 0
	start, end := 0, 0
	flag := false
	for k, v := range S {
		if string(v) == "(" {
			num++
			if !flag {
				start = k
				flag = true
			}
		} else {
			num--
			if num == 0 {
				end = k
				flag = false
				res += S[start+1 : end]
				start = end
			}
		}
	}
	return res
}

func removeOuterParentheses2(S string) string {
	var res string
	opened := 0
	for c := range S {
		o := opened + 1
		if c == '(' && o > 0 {
			res = fmt.Sprintf(res, c)
		}
		h := opened - 1
		if c == ')' && h > 1 {
			res = fmt.Sprintf(res,c)
		}
	}
	return res
}
