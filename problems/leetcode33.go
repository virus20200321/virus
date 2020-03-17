package problems

func search(nums []int, target int) int {
	len := len(nums)
	var l, r = 0, len - 1
	if len == 0 {
		return -1
	}
	if len == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	i := findRotateIndex(nums)
	if i == -1 {
		l, r = 0, len-1
	} else {
		if target > nums[0] {
			r = i - 1
		} else if target < nums[0] {
			l = i
		} else {
			return 0
		}
	}

	//fmt.Println("search range:", l, r)
	return binarySearch(nums, target, l, r)
}

func binarySearch(nums []int, target int, l, r int) int {
	if l > r {
		return -1
	}
	p := (l + r) / 2
	if nums[p] == target {
		return p
	} else if nums[p] > target {
		return binarySearch(nums, target, l, p-1)
	} else {
		return binarySearch(nums, target, p+1, r)

	}
}
func findRotateIndex(nums []int) int {
	var left, right = 0, len(nums) - 1
	for left < right {
		p := (left + right) / 2
		if nums[p] > nums[p+1] {
			return p + 1
		}
		if nums[p] > nums[left] {
			left = p
		} else {
			right = p
		}

	}
	return -1
}
