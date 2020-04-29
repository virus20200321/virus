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

func getIndex(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			return mid + 1
		}
		if nums[mid] > nums[l] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

func bSearch(nums []int, t int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		if nums[mid] == t {
			return mid
		} else if nums[mid] > t {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

func search(nums []int, target int) int {
	if pivot := getIndex(nums); pivot == -1 {
		return bSearch(nums, target)
	} else if pivot == target {
		return pivot
	} else if target < nums[0] {
		return bSearch(nums[:pivot], target)
	} else {
		return bSearch(nums[pivot:], target)
	}

}
