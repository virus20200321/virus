package problems

func findDuplicate(nums []int) int {
	if len(nums) <= 2 {
		return nums[0]
	}
	l, r := 1, len(nums)-1

	for l < r {
		mid := (l + r) / 2

		cnt := 0

		for _, num := range nums {
			if num <= mid {
				cnt++
			}
		}
		if cnt <= mid {
			l = mid + 1
		} else {
			r = mid
		}
		//fmt.Printf("range %d %d mid=%d cnt = %d\n", l, r, mid, cnt)
	}
	return l
}
