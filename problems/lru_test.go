package problems

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"testing"
)

func TestLruCache(t *testing.T) {
	logrus.SetLevel(logrus.DebugLevel)

	lru := NewLruCache(5)
	for i := 0; i < 5; i++ {

		lru.Push(cast.ToString(i), cast.ToString(i))
	}
	lru.Print()

	// get
	lru.Get("2")
	lru.Print()

	// push when list is full
	lru.Push("6", "6")
	lru.Print()

}
