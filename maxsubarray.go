package NewString

func MaxSumSubarray(nums []int, window int) int {
	start, end := 0, 0
	sum := 0
	maxsum := 0
	for end < len(nums) {
		sum += nums[end]
		if end-start+1 == window {
			if maxsum < sum {
				maxsum = sum
			}
			sum -= nums[start]
			start++
		}
		end++
	}
	return maxsum
}
