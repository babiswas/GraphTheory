package NewString

func Longest_substring_no_repeat(str string) int {
	left, right := 0, 0
	mp := make(map[string]int)
	max_length := 0

	for i := 0; i < len(str); i++ {
		mp[string(str[i])] = -1
	}

	for right < len(str) {
		if mp[string(str[right])] == -1 {
			mp[string(str[right])] = right
			if (right - left + 1) > max_length {
				max_length = right - left + 1
			}
			right += 1
		} else if mp[string(str[right])] != -1 {
			if mp[string(str[right])] < left {
				mp[string(str[right])] = right
				if (right - left + 1) > max_length {
					max_length = right - left + 1
				}
				right += 1
			} else {
				left = mp[string(str[right])] + 1
				mp[string(str[right])] = right
				if (right - left + 1) > max_length {
					max_length = right - left + 1
				}
				right += 1
			}
		}
	}
	return max_length
}
