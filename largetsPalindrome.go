package NewString

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(str string) string {
	test := strings.Split(str, "")
	start_index := 0
	end_index := len(str) - 1
	for start_index < end_index {
		test[start_index], test[end_index] = test[end_index], test[start_index]
		start_index += 1
		end_index -= 1
	}
	return strings.Join(test, "")
}

func FindNextLargest(str []string) string {
	count := make([]int, 10)
	for i := 0; i < 10; i++ {
		count[i] = 0
	}
	pos := 0
	digit := 0
	for pos = len(str) - 1; pos >= 0; pos-- {
		num, _ := strconv.Atoi(str[pos])
		count[num] += 1
		if num < digit {
			digit = num
			break
		}
		digit = num
	}

	if pos == -1 {
		return "Not Possible"
	}
	fmt.Println("POSITION:", digit)
	fmt.Println(count)

	for value := digit + 1; value < 10; value++ {
		if count[value] > 0 {
			count[value]--
			str[pos] = strconv.Itoa(value)
			pos += 1
			break
		}
	}

	for val := 0; val < 10; val++ {
		for count[val] > 0 {
			str[pos] = strconv.Itoa(val)
			pos += 1
			count[val]--
		}
	}
	return strings.Join(str, "")
}

func NextLargestPalindrome(num int) string {
	temp_str := strconv.Itoa(num)
	ch_arr := strings.Split(temp_str, "")
	mid := ""
	index := -1
	if len(temp_str)%2 == 0 {
		index = len(temp_str) / 2
		mid = ""
	} else {
		index := len(temp_str) / 2
		mid = string(temp_str[index])
	}
	sub_str := FindNextLargest(ch_arr[0:index])
	if sub_str == "Not Possible" {
		return sub_str
	}
	return sub_str + mid + reverse(sub_str)
}
