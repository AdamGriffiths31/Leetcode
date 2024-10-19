func compress(chars []byte) int {
	i := 0
	j := 0

	for j < len(chars) {
		char := chars[j]
		count := 0

		for j < len(chars) && chars[j] == char {
			j++
			count++
		}

		chars[i] = char
		i++

		if count > 1 {
			for _, c := range []byte(fmt.Sprintf("%d", count)) {
				chars[i] = c
				i++
			}
		}
	}

	return i
}
