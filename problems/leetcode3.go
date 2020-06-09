func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {

		}
	}
}