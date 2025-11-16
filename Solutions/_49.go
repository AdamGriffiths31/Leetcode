func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]byte][]string)
	for _, s := range strs {
		groups[strKey(s)] = append(groups[strKey(s)], s)
	}
	result := make([][]string, 0, len(groups))
	for _, v := range groups {
		result = append(result, v)
	}
	return result
}

func strKey(v string) [26]byte {
	res := [26]byte{}
	for i := range v {
		res[v[i]-'a']++
	}
	return res
}
