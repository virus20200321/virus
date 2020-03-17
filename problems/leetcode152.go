package problems

func maxProduct(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {

		tmp := dp[i-1] * nums[i]
		if tmp > 0 {
			dp[i] = tmp
		} else {
			dp[i] = nums[i]
			if i >= 2 {
				try := nums[i] * dp[i-1] * dp[i-2]
				if dp[i] < try {
					dp[i] = try
				}
			}
		}
		if max < dp[i] {
			max = dp[i]
		}
	}

	return max
}
