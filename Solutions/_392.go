func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	sPointer := 0

	for tPointer := range t {
		if t[tPointer] == s[sPointer] {
			sPointer++
		}
		if sPointer >= len(s) {
			return true
		}
	}

	return false
}