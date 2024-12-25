package easy

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 0; i < len(strs); i++ {
		for j := 0; j < len(prefix) && j < len(strs[i]); j++ {
			if prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				break
			}
		}
		if len(prefix) > len(strs[i]) {
			prefix = prefix[:len(strs[i])]
		}
	}
	return prefix
}