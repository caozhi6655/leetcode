package easy

//MaxSubArray 最大子序列
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		panic("nil")
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	}

	max := dp[0]
	for _, v := range dp {
		if v > max {
			max = v
		}
	}

	return max
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
