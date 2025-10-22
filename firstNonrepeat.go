package NewString

import (
	"strconv"
	"strings"
)

func FirstNonRepeating(str string) string {
	if len(str) == 0 {
		return ""
	}
	str = strings.ToLower(str)
	status := false
	for i := 0; i < len(str); i++ {
		status = false
		for j := 0; j < len(str); j++ {
			if str[i] == str[j] && i != j {
				status = true
			}
		}
		if !status {
			return string(str[i])
		}
	}
	return ""
}

func FirstNonRepeatingV2(str string) string {
	if len(str) == 0 {
		return ""
	}
	first_nonrepeat, idx := "", len(str)+1
	mp := make(map[rune]string)
	str = strings.ToLower(str)
	for index, ch := range str {
		_, ok := mp[ch]
		if ok {
			num := "," + strconv.Itoa(index)
			mp[ch] += num
		} else {
			mp[ch] = strconv.Itoa(index)
		}
	}
	for key, val := range mp {
		val := strings.Split(val, ",")

		if len(val) == 1 {
			nm, _ := strconv.Atoi(val[0])
			if idx > nm {
				first_nonrepeat = string(key)
				idx = nm
			}
		}
	}
	return first_nonrepeat
}

func FirstNonRepeatingCharV3(str string) string {
	mp := make(map[rune]int)
	start_index := len(str) + 1
	char := ""
	for index, ch := range str {
		_, ok := mp[ch]
		if ok {
			mp[ch] = -2
		} else {
			mp[ch] = index
		}
	}
	for key, val := range mp {
		if val > 0 {
			if start_index > val {
				start_index = val
				char = string(key)
			}
		}
	}
	return char
}
