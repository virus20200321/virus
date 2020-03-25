package review

import "fmt"

type point []int

func (p point) x() int {
	return p[0]
}
func (p point) y() int {
	return p[1]
}

func maxPoints(points [][]int) int {
	atSamePoint := func(p1, p2 point) bool {
		return p1.x() == p2.x() && p1.y() == p2.y()

	}

	gcdFunc := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	retMax := 0
	for i := 0; i < len(points); i++ {
		kDict := make(map[string]int)
		dup := 0
		tmpMax := 0

		for j := 0; j < len(points); j++ {
			if i == j {
				continue
			}
			if atSamePoint(points[i], points[j]) {
				dup++
			} else {
				up := point(points[i]).y() - point(points[j]).y()
				down := point(points[i]).x() - point(points[j]).x()
				gcd := gcdFunc(up, down)
				if gcd != 0 {
					up, down = up/gcd, down/gcd
				}
				key := fmt.Sprintf("%d/%d", up, down)
				if _, ok := kDict[key]; ok {
					kDict[key] += 1
				} else {
					kDict[key] = 1 // exclude dup and points[i] itself
				}
				if tmpMax < kDict[key] {
					tmpMax = kDict[key]
				}

			}

		}
		if dup+1+tmpMax > retMax {
			retMax = dup + 1 + tmpMax
		}
	}
	return retMax
}

func longestPalindrome(s string) string {
	l := len(s)
	if l < 2 {
		return s
	}
	max := 1
	ret := s[0:1]
	dp := make([][]bool, l)
	for i, _ := range dp {
		dp[i] = make([]bool, l)
	}

	for r := 1; r < l; r++ { //!! 保证 dp[r-1][l+1] 有值
		for l := 0; l < r; l++ {
			if r-l < 3 { //!! <=3 就是4个值了
				dp[r][l] = s[r] == s[l]
			} else {
				dp[r][l] = dp[r-1][l+1] && s[r] == s[l]
			}
			if dp[r][l] {
				if r-l+1 > max {
					max = r - l + 1
					ret = s[l : r+1]
				}
			}
		}
	}
	return ret
}

func search(nums []int, target int) int {
	rotate := findRotateIndex(nums)
	if rotate == -1 {
		return binarySearch(nums, target)
	}
	i := -1
	if nums[rotate] == target {
		return rotate
	} else if target >= nums[0] {
		fmt.Println("search in", nums[:rotate])
		i = binarySearch(nums[:rotate], target)

	} else {
		fmt.Println("search in", nums[rotate:])

		i = binarySearch(nums[rotate:], target)
		if i != -1 {
			i = rotate + i
		}
	}

	return i

}

func binarySearch(nums []int, t int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		p := (left + right) / 2
		if nums[p] == t {
			return p
		} else if nums[p] > t {
			right = p - 1
		} else {
			left = p + 1
		}

	}
	return -1
}

func findRotateIndex(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		p := (left + right) / 2
		if nums[p] > nums[p+1] {
			return p + 1
		}
		if nums[p] > nums[left] {
			//TODO is this ok? it is !
			left = p + 1
		} else {
			right = p
		}

	}
	return -1
}
