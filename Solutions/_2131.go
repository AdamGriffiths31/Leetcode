func longestPalindrome(words []string) int {
	hashMap := make(map[string]int)
	res := 0
	for _, word := range words {
		if _, ok := hashMap[word]; ok {
			res += 4
			hashMap[word]--
			if hashMap[word] <= 0 {
				delete(hashMap, word)
			}
		} else {
			revWord := string([]byte{word[1], word[0]})
			hashMap[revWord]++
		}
	}

	for word := range hashMap {
		if word[0] == word[1] {
			res += 2
			break
		}
	}

	return res
}