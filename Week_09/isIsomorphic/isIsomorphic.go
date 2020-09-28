package isIsomorphic

func isIsomorphic(s string, t string) bool {
	return isMorphic(s, t) && isMorphic(t, s)
}
func isMorphic(s, t string) bool {
	m, n := map[byte]byte{}, len(s)
	for i := 0; i < n; i++ {
		c1, c2 := s[i], t[i]
		if c, ok := m[c1]; ok {
			if c != c2 {
				return false
			}
		} else {
			m[c1] = c2
		}
	}
	return true
}
