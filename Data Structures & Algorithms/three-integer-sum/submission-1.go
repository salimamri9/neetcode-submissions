import (
	"slices"
)

func threeSum(nums []int) [][]int {
	slices.Sort(nums)
	var Result [][]int
	for i, num := range nums {
		if num > 0 {
			break
		} else if i > 0 && num == nums[i-1] {
			continue
		}
		for j, k := i+1, len(nums)-1; j < k; {
			if num+nums[j]+nums[k] == 0 {
				Result = append(Result, []int{num, nums[j], nums[k]})
				// you have to increment j or decrement k to avoid infinite loop and move forward with the algorithm
				// but we do both to avoid uncessary iterations since we know that the next number will trigger an if statement and
				// we can skip it since we know
				// that nums[i] + nums[j] + nums[k] = 0, then it is mathematically impossible for nums[i] + (new_j) + nums[k] to equal zero unless
				// new_j is the same value as the old j.
				j++
				k--
                for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if num+nums[j]+nums[k] < 0 {
				j++
			} else if num+nums[j]+nums[k] > 0 {
				k--
			}
		}
	}

	return Result
}

