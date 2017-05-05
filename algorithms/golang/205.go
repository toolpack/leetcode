package golang

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m1 := make(map[byte]byte)
	m2 := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if d, ok := m1[s[i]]; ok {
			if d != t[i] {
				return false
			}
		} else {
			m1[s[i]] = t[i]
		}
		if d, ok := m2[t[i]]; ok {
			if d != s[i] {
				return false
			}
		} else {
			m2[t[i]] = s[i]
		}
	}
	return true
}
