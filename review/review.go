package review

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dup := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			dup++
		} else {
			nums[i-dup] = nums[i]
		}
	}
	return len(nums) - dup
}
