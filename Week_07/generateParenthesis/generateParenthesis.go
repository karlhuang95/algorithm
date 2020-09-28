package generateParenthesis

func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}

	combs := []string{"("}
	for i := 1; i < n*2; i++ {
		for l := len(combs); l > 0; l-- {
			// combs[l-1]中的左括号与右括号数量
			var (
				lp, rp int
			)
			for _, p := range combs[l-1] {
				switch p {
				case '(':
					lp++
				case ')':
					rp++
				default:
					panic("invalid character: " + string(p))
				}
			}

			// 确定可能的组合
			if lp-rp > 0 {
				if lp < n {
					combs = append(combs, combs[l-1]+")")
					combs[l-1] += "("
				} else {
					combs[l-1] += ")"
				}
			} else {
				combs[l-1] += "("
			}
		}
	}

	return combs
}
