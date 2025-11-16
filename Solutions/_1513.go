func numSub(s string) int {
	consec := 0
	res := 0
	for _, v := range s {
		if string(v) == "1" {
			consec++
		} else {
			consec = 0
		}
		res += consec
	}
	return res % 1_000_000_007
}
