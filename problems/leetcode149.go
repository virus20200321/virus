package problems

import (
	"fmt"
)

func Gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func equal(p1, p2 []int) bool {
	return p1[0] == p2[0] && p1[1] == p2[1]
}
func badCase(points [][]int) int {
	if len(points) < 3 {
		return len(points)
	}
	eqCnt := 0
	for i := 0; i < len(points)-1; i++ {
		if equal(points[i], points[i+1]) {
			eqCnt++
		} else {
			break
		}
	}
	if eqCnt == len(points)-1 {
		return len(points)
	}
	return -1
}

func MaxPoints(points [][]int) int {
	if cnt := badCase(points); cnt > 0 {
		return cnt
	}

	length := len(points)
	cnt := 0
	max := 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if i == j || equal(points[i], points[j]) {
				continue
			}
			son := points[i][0] - points[j][0]
			mother := points[i][1] - points[j][1]
			if son != 0 && mother != 0 {
				gcd := Gcd(son, mother)
				son = son / gcd
				mother = mother / gcd
			}
			fmt.Println("line", points[i], points[j])
			cnt = 2 // already has 2 points
			if cnt > max {
				max = cnt
			}
			for k := 0; k < length; k++ {
				if k == i || k == j {
					continue
				}

				if son*points[k][1]-mother*points[k][0] == son*points[i][1]-mother*points[i][0] {
					cnt++
					if cnt > max {
						fmt.Printf("update max from %d to %d\n", max, cnt)

						max = cnt
					}
				}

			}
			cnt = 0
		}
	}

	return max
}

func MaxPoints2(points [][]int) int {
	if len(points) < 3 {
		return len(points)
	}

	length := len(points)
	res := 0
	for i := 0; i < length; i++ {
		kMap := make(map[string]int)
		dup := 0
		iMax := 0
		for j := 0; j < length; j++ {
			if j == i {
				continue
			}
			if equal(points[i], points[j]) {
				dup++
			}
			up := points[i][1] - points[j][1]
			down := points[i][0] - points[j][0]
			gcd := Gcd(up, down)
			if gcd != 0 {
				up = up / gcd
				down = down / gcd
			}
			key := fmt.Sprintf("%d/%d", up, down)
			if cnt, ok := kMap[key]; ok {
				kMap[key] = cnt + 1
			} else {
				kMap[key] = 1
			}

			if iMax < kMap[key] {
				iMax = kMap[key]
			}
		}
		if res < iMax+dup+1 {
			res = iMax + dup + 1
		}
	}
	return res
}
