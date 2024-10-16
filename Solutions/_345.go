var vowels = map[byte]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
}

func reverseVowels(s string) string {
	left := 0
	right := len(s) - 1
	bytes := []byte(s)
	inputLower := strings.ToLower(s)

	for left <= right {
		_, okLeft := vowels[inputLower[left]]
		if !okLeft {
			left++
		}

		_, okRight := vowels[inputLower[right]]
		if !okRight {
			right--
		}

		if okLeft && okRight {
			bytes[left], bytes[right] = s[right], s[left]
			left++
			right--
		}
	}

	return string(bytes)
}
