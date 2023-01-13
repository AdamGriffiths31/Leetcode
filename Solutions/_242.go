func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	hashmapS := make(map[string]int)
	hashmapT := make(map[string]int)

	for i := range s {
		hashmapS[string(s[i])]++
		hashmapT[string(t[i])]++
	}

	for v := range hashmapS {
		if hashmapT[v] == 0 {
			return false
		}

		if hashmapT[v] != hashmapS[v] {
			return false
		}
	}

	return true
}
