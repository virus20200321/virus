package problems

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/spf13/cast"
	"regexp"
	"testing"
)

func TestGcd(t *testing.T) {

	var x,y int = 9,6

	assert.Equal(t,Gcd(x,y),int(3))

	x,y = 192,123
	assert.Equal(t,Gcd(x,y),int(3))

	x,y = 60,15
	assert.Equal(t,Gcd(x,y),int(15))

	x,y = -5,-14
	assert.Equal(t,Gcd(x,y),int(-1))

}

func TestMaxPoints(t *testing.T) {
	var re = regexp.MustCompile(`(?m)(\d+)`)
	var str = `[[0,-12],[5,2],[2,5],[0,-5],[1,5],[2,-2],[5,-4],[3,4],[-2,4],[-1,4],[0,-5],[0,-8],[-2,-1],[0,-11],[0,-9]]`
	data := make([][]int,0)

	ints := make([]int,2)
	for i, match := range re.FindAllString(str, -1) {
		//fmt.Println(match, "found at index", i)
		ints[i%2] = cast.ToInt(match)
		if i % 2 == 1{
			temp := make([]int,2)
			copy(temp, ints)
			data = append(data, temp)
		}
	}
	fmt.Println(data)

	max := MaxPoints(data)
	fmt.Println(max)
}