package review

import (
	"fmt"
	"testing"
)

func Test_findRotate(t *testing.T) {
	i := findRotateIndex([]int{4, 5, 6, 7, 1, 2, 3})
	fmt.Println(i)
	i = findRotateIndex([]int{1, 2, 3, 4, 5})
	fmt.Println(i)

}

func Test_search(t *testing.T) {
	i := search([]int{4, 5, 6, 7, 1, 2, 3}, 3)
	fmt.Println(i)
}
