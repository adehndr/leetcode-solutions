package twopointers

func IsSubsequence(s string, t string) bool {
	i := 0
	j := 0
	for j < len(t) && i < len(s) {
		if string(s[i]) == string(t[j]) {
			i += 1
		}
		j += 1
	}
	return i >= len(s)
}
