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

func findDuplicate2(nums []int) int {
	//nums range 1~n len = n+1
	//sorted range 1~n len = n

	if len(nums) <= 2 {
		return nums[0]
	}
	l, r := 1, len(nums)-1

	for l <= r {
		cnt := 0
		mid := (l + r) / 2
		for i := 0; i < len(nums); i++ {
			if nums[i] <= mid {
				cnt++
			}
		}
		if cnt > len(nums)/2 {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
