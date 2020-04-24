package contents

import (
	"fmt"
	"leetcode213/reminder/pkg/macros"
	"math/rand"
	"time"
)

//Probability arises with failure times
var CanLuFailure = float64(0)

func CanLu() (msg string, can bool) {
	curProbability := macros.MinProbability + CanLuFailure*0.3
	if curProbability >= macros.MaxProbability {
		curProbability = macros.MaxProbability
	}
	rand.Seed(time.Now().UnixNano())

	num := rand.Float64()
	ok := num <= curProbability
	if !ok {
		CanLuFailure++
	} else {
		CanLuFailure = 0
	}
	return fmt.Sprintf("current result: %0.2f\nfailure times: %1.0f", num, CanLuFailure), ok
}
