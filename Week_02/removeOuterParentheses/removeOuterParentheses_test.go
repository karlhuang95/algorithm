package removeOuterParentheses

import "testing"

func TestRemoveOuterParemtheses(t *testing.T)  {
	S := "(()())(())"
	info := removeOuterParentheses(S)
	t.Log(info)
}
