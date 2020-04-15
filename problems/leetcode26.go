package problems

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dupCnt := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			dupCnt++
		} else {
			nums[i-dupCnt] = nums[i]
		}
	}

	return len(nums) - dupCnt
}
