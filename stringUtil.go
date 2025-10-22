package NewString

import "strings"

func FindFirstRepeatChar(str string) string {
	str = strings.ToLower(str)
	mp := make(map[rune]int)
	temp := []rune(str)
	for index, ch := range temp {
		_, ok := mp[ch]
		if ok {
			return string(ch)
		}
		mp[ch] = index
	}
	return ""
}

func FindFirstRepeatCharV2(str string) string {
	str = strings.ToLower(str)
	for i := 0; i < len(str); i++ {
		for j := 0; j < i; j++ {
			if str[i] == str[j] {
				return string(str[i])
			}
		}
	}
	return ""
}
