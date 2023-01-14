func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	lPointer := 0
	rPointer := len(s) - 1

	for lPointer < rPointer {
		for !isAlphaNumeric(s[lPointer]) && lPointer < rPointer {
			lPointer++
		}

		for !isAlphaNumeric(s[rPointer]) && lPointer < rPointer {
			rPointer--
		}

		fmt.Printf("%v - %v", lPointer, rPointer)
		if s[lPointer] != s[rPointer] {
			return false
		}

		lPointer++
		rPointer--
	}
	return true
}

func isAlphaNumeric(c byte) bool {

	if (c >= 'a' && c <= 'z') || c >= 'A' && c <= 'Z' || (c >= '0' && c <= '9') {
		return true
	}

	return false
}
