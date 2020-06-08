package problems

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ret := make([][]int, 0)
	// fmt.Println(nums)
	for i := 0; i < len(nums)-2; i++ {
		sum := 0 - nums[i]
		j := i + 1
		k := len(nums) - 1
		// fmt.Println("j,k",j,k)
		for j < k {
			if nums[j]+nums[k] == sum {
				if j == i+1 || (j > i+1 && nums[j] != nums[j-1]) {
					ret = append(ret, []int{nums[i], nums[j], nums[k]})

				}
				j++
				k--
			} else if nums[j]+nums[k] < sum {
				j++
			} else {
				k--
			}

		}

		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return ret
}
