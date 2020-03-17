package problems

import (
	"fmt"
	"testing"
)

func Test_longestPalindrome(t *testing.T) {
	src := "cvabbac"
	r := longestPalindrome2(src)
	fmt.Println(r)
}
