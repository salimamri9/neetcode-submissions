

func maxSubArray(nums []int) int {

	globalMax := nums[0]
	currentMax := nums[0]
	n := len(nums)

	for i := 1; i < n; i++ {
		if currentMax < 0 {
			currentMax = nums[i]
		} else {
			currentMax += nums[i]
		}

		if globalMax < currentMax {
			globalMax = currentMax
		}

	}

	return globalMax
}