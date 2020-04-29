package problems

import "fmt"

func getIndex(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] > nums[mid+1] {
			return mid + 1
		}
		if nums[mid] > nums[l] {
			l = mid
		} else {
			r = mid
		}
	}
	return -1
}

func bSearch(nums []int, t int) int {
	l, r := 0, len(nums)-1
	fmt.Println("search from", l, " - ", r, nums)
	for l <= r {
		mid := (l + r) / 2
		fmt.Println("mid=", mid, nums[mid])
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
	pivot := getIndex(nums)
	fmt.Println("pivot=", pivot)
	if pivot == -1 {
		return bSearch(nums, target)
	} else if nums[pivot] == target {
		return pivot
	}

	if target >= nums[0] {
		return bSearch(nums[:pivot], target)
	} else {
		if ret := bSearch(nums[pivot:], target); ret != -1 {
			return pivot + ret
		}
		return -1
	}
}
