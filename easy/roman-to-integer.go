package easy

func romanToInt(s string) int {
	res := 0
	converter := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := len(s) - 1; i >= 0; i-- {
		if i < len(s)-1 && converter[s[i+1]] > converter[s[i]] {
			res -= converter[s[i]]
		} else {
			res += converter[s[i]]
		}
	}
	return res
}