package problems

import (
	"fmt"
	"testing"
)

func Test_findRotate(t *testing.T) {
	i := findRotateIndex([]int{5, 6, 7, 8, 1, 2, 3})
	fmt.Println("rotate index = ", i)
	i = findRotateIndex([]int{5, 1, 2})
	fmt.Println("rotate index = ", i)
	i = findRotateIndex([]int{1, 3})
	fmt.Println("rotate index = ", i)

}

func Test_search(t *testing.T) {
	i := search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	fmt.Println("result index = ", i)
	i = search([]int{1, 3}, 1)
	fmt.Println("result index = ", i)

}
