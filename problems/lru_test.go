package problems

import (
	"github.com/fatih/color"
	"testing"
)

func TestLruCache(t *testing.T) {

	lruCache := Constructor(2)
	r := lruCache.Get(1)
	color.Yellow("r=%d lruCache: %s", r, &lruCache)

	lruCache.Put(1, 1)
	color.Yellow("r=%d lruCache: %s", r, &lruCache)
	lruCache.Put(3, 3)
	color.Yellow("r=%d lruCache: %s", r, &lruCache)

	lruCache.Put(6, 6)
	color.Yellow("r=%d lruCache: %s", r, &lruCache)

	r = lruCache.Get(3)
	color.Yellow("r=%d lruCache: %s", r, &lruCache)

}
