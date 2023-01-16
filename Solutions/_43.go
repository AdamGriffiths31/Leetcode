func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	num1 = reverseString(num1)
	num2 = reverseString(num2)

	res := make([]int, len(num1)+len(num2))

	for i1 := 0; i1 < len(num1); i1++ {
		for i2 := 0; i2 < len(num2); i2++ {
			i1val := int(num1[i1]) - 48
			i2val := int(num2[i2]) - 48
			var digit int = i1val * i2val
			res[i1+i2] += digit
			res[i1+i2+1] += res[i1+i2] / 10
			res[i1+i2] = res[i1+i2] % 10
		}
	}
	fmt.Printf("%v\n", res)
	res = reverseSlice(res)

	i := 0
	for i < len(res) && res[i] == 0 {
		i++
	}

	res = res[i:]

	result := ""

	for _, v := range res {
		fmt.Printf("%v\n", v)
		result += strconv.Itoa(v)

	}
	return result
}

func reverseString(num string) string {
	var s string
	for _, v := range num {
		s = string(v) + s
	}

	return s
}

func reverseSlice(slice []int) []int {
	rev := make([]int, 0)
	for i := range slice {
		rev = append(rev, slice[len(slice)-1-i])
	}

	return rev
}
