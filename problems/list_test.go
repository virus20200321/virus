package problems

import (
	"github.com/spf13/cast"
	"testing"
)

func Test_DoubleLink(t *testing.T) {
	dl := DoubleList{
		Head: nil,
		Tail: nil,
	}

	for i := 0; i < 50; i++ {
		dl.Push(cast.ToString(i))
	}
	dl.Delete("22")

	dl.Print()
	dl.ReversePrint()
}
