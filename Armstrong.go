package main

import (
	"math"
)

// Armstrong Number
func checkArmstrong(num int64) bool {
	tempNum := num
	var sum int64

	for tempNum > 0 {
		sum = sum + int64(math.Pow(float64(tempNum%10), 3))
		// println(tempNum, sum)
		tempNum = tempNum / 10
	}

	if sum == num {
		return true
	}
	return false
}
