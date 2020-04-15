package problems

// 3 3 4 : 1 2 1
func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	cnt := 1
	ret := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == ret {
			cnt++
		} else {
			cnt--
		}
		if cnt == 0 {
			ret = nums[i+1]
		}
	}
	return ret
}
