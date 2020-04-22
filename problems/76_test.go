package problems

import (
	"fmt"
	"testing"
)

func Test_minWindow(t *testing.T) {
	ret := minWindow("aaaaaaaaaaaabbbbbcdd", "abcdd")
	fmt.Println(ret)
}
