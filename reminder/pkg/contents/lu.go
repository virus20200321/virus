package contents

import (
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
	rand.Float64()
	return "", false
}
