package NewString

import "strings"

func FindIfStringsMatch(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)
	return strings.Contains(str1+str1, str2)
}

func KMP(hay, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	prevlps, i := 0, 1
	lps := make([]int, len(needle))
	for i := 0; i < len(needle); i++ {
		lps[i] = 0
	}
	for i < len(needle) {
		if needle[i] == needle[prevlps] {
			lps[i] = prevlps + 1
			prevlps = prevlps + 1
			i = i + 1
		} else if prevlps == 0 {
			lps[i] = 0
			i = i + 1
		} else {
			prevlps = lps[prevlps-1]
		}
	}
	i, j := 0, 0
	for i < len(hay) {
		if hay[i] == needle[j] {
			i, j = i+1, j+1
		} else {
			if j == 0 {
				i = i + 1
			} else {
				j = lps[j-1]
			}
		}
		if j == len(needle) {
			return i - len(needle)
		}
	}
	return -1
}
