func mergeAlternately(word1 string, word2 string) string {
	i := 0
	j := 0
	result := ""

	for i < len(word1) && j < len(word2) {
		result = result + string(word1[i])
		result = result + string(word2[j])
		i++
		j++
	}

	result += word1[i:]
	result += word2[j:]

	return result
}
