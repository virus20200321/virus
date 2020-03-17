package problems

import (
	"fmt"
	"testing"
)

func Test_findDuplicate(t *testing.T) {

	ret := findDuplicate([]int{1, 3, 4, 2, 2})
	fmt.Println(ret)
	ret = findDuplicate([]int{3, 1, 3, 4, 2})
	fmt.Println(ret)

}
