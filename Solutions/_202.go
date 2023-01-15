func isHappy(n int) bool {
	if n <= 1 || n == 4 {
		return n == 1
	}

	numStr := strconv.Itoa(n)
	total := 0
	for _, digit := range numStr {
		value, _ := strconv.Atoi(string(digit))
		total = total + (value * value)
	}

	fmt.Println(total)

	return isHappy(total)
}