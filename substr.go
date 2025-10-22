package NewString

import (
	"fmt"
	"strings"
)

func FindKAtmost(str1 string, K int) int {
	result := 0
	start_index := 0
	end_index := 0
	temp_map := make(map[string]int)
	for end_index < len(str1) {
		_, ok := temp_map[string(str1[end_index])]
		if !ok {
			temp_map[string(str1[end_index])] = 1
		} else {
			temp_map[string(str1[end_index])] += 1
		}
		if len(temp_map) <= K {
			result += (end_index - start_index) + 1
		} else if len(temp_map) > K {
			for len(temp_map) > K {
				ch := string(str1[start_index])
				start_index += 1
				temp_map[ch] -= 1
				if temp_map[ch] == 0 {
					delete(temp_map, ch)
				}
				if len(temp_map) == K {
					result += end_index - start_index + 1
					break
				}
			}
		}
		end_index += 1
	}
	return result
}

func FindDistinctK(str1 string, K int) int {
	if K-1 != 0 {
		return FindKAtmost(str1, K) - FindKAtmost(str1, K-1)
	} else {
		return FindKAtmost(str1, K)
	}

}

func CountSubstringKchar(str1 string, K int) int {
	str1 = strings.ToLower(str1)
	list_of_substr := make([]string, 0)
	mp := make(map[string]bool)
	temp := ""
	for i := 0; i < len(str1); i++ {
		temp = temp + string(str1[i])
		list_of_substr = append(list_of_substr, temp)
		for j := i + 1; j < len(str1); j++ {
			temp += string(str1[j])
			list_of_substr = append(list_of_substr, temp)
		}
		temp = ""
	}
	final_count := 0
	count := 0
	fmt.Println("The list of substring is:", list_of_substr)
	for _, val := range list_of_substr {
		for _, ch := range val {
			_, ok := mp[string(ch)]
			if !ok {
				count = count + 1
				mp[string(ch)] = true
			}
		}
		if count == K {
			final_count += 1
		}
		count = 0
		clear(mp)
	}
	return final_count
}
