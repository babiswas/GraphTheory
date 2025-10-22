package NewString

import (
	"fmt"
	"math"
)

func MinimumWindow(str string, substr string) (int, int, int) {
	temp_lngth := 0
	start, end := 0, 0
	max_lngth := math.MaxInt64
	mp := make(map[string]int)
	for i := 0; i < len(substr); i++ {
		_, ok := mp[string(substr[i])]
		if !ok {
			mp[string(substr[i])] = 1
		} else {
			mp[string(substr[i])] += 1
		}
	}
	tp := make(map[string]int)
	for i := 0; i < len(substr); i++ {
		_, ok := tp[string(substr[i])]
		if !ok {
			tp[string(substr[i])] = 0
		}
	}

	isMatching := func(m1, m2 map[string]int) bool {
		distinct := 0
		required := len(m1)
		for key, values := range m1 {
			fmt.Println(m2)
			if m2[key] >= values {
				distinct += 1
			}
		}
		return distinct == required
	}

	startIndex, endIndex := 0, 0
	for endIndex < len(str) {
		_, ok := tp[string(str[endIndex])]
		if ok {
			tp[string(str[endIndex])] += 1
		}
		if isMatching(mp, tp) {
			temp_lngth = endIndex - startIndex + 1
			if temp_lngth < max_lngth {
				max_lngth = temp_lngth
				start, end = startIndex, endIndex
			}
			for isMatching(mp, tp) {
				temp_lngth = endIndex - startIndex + 1
				if temp_lngth < max_lngth {
					max_lngth = temp_lngth
					start, end = startIndex, endIndex
				}
				_, ok := tp[string(str[startIndex])]
				if ok {
					if tp[string(str[startIndex])] > 0 {
						tp[string(str[startIndex])] -= 1
					}
				}
				startIndex += 1
			}

		}
		endIndex += 1
	}
	return start, end, max_lngth
}
