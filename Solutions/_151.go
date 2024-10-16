func reverseWords(s string) string {
	result := ""
	start := 0
	end := 0

	for end < len(s) {
		if s[end] != ' ' {
			start = end
			for end < len(s) && s[end] != ' ' {
				end++
			}
			result = string(s[start:end]) + " " + result
			end -= 1
		}
		end += 1
	}

	return strings.TrimRight(result, " ")
}
