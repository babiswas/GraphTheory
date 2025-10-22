package NewString

import (
	"strconv"
)

func CompressString(str []string) int {
	index := 0
	i := 0
	current_char := string(str[i])
	count := 0
	for i < len(str) {
		if string(str[i]) == current_char {
			count += 1
		} else if string(str[i]) != current_char {
			temp_str := str[i]
			if count == 1 {
				index += 1
				str[index] = temp_str
				current_char = temp_str
				count = 1
			} else {
				temp := strconv.Itoa(count)
				if len(temp) == 1 {
					index += 1
					str[index] = string(temp)
					index += 1
					str[index] = temp_str
					current_char = temp_str
					count = 1
				} else {
					for _, val := range temp {
						index += 1
						str[index] = string(val)
					}
					index += 1
					str[index] = temp_str
					current_char = temp_str
					count = 1
				}
			}
		}
		i += 1
	}

	if count > 1 {
		temp := strconv.Itoa(count)
		if len(temp) == 1 {
			index += 1
			str[index] = temp
			return index + 1
		} else {
			for _, val := range temp {
				index += 1
				str[index] = string(val)
			}
			return index + 1
		}
	}

	return index + 1
}
