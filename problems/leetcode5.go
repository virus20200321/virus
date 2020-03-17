package problems

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	dp := make([][]bool, len(s))
	for i, _ := range dp {
		dp[i] = make([]bool, len(s))
	}
	//init
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}
	max := 1
	var res string = s[:1]
	//fmt.Println(dp)

	for rIndex := 1; rIndex < len(s); rIndex++ {
		for lIndex := 0; lIndex < rIndex; lIndex++ {
			if s[rIndex] == s[lIndex] {
				if rIndex-lIndex < 3 {
					dp[lIndex][rIndex] = true
				} else {
					dp[lIndex][rIndex] = dp[lIndex+1][rIndex-1]
				}
			} else {
				dp[lIndex][rIndex] = false
			}

			if dp[lIndex][rIndex] {
				if max < rIndex-lIndex+1 {
					max = rIndex - lIndex + 1
					res = s[lIndex : rIndex+1]
				}
			}
		}
	}
	return res
}

func longestPalindrome2(s string) string {
	if len(s) < 2 {
		return s
	}
	max := 1
	res := ""
	for i := 1; i < len(s); i++ {
		for j := 0; j < i; j++ {
			if isPalindrome(s[j : i+1]) {
				if i-j+1 > max {
					max = i - j + 1
					res = s[j : i+1]
				}
			}

		}
	}

	return res
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1

	for i <= j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
