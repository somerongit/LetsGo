package main

import (
	"math"
)

// Armstrong Number
/*
A number is called an Armstrong number if the sum of the cubes
of the digits of the number is equal to the number.
E.g 153 = 1^3 + 5^3 + 3^3.
*/

func checkArmstrong(num int64) (check bool) {
	tempNum := num
	var sum int64
	check = false

	for tempNum > 0 {
		sum = sum + int64(math.Pow(float64(tempNum%10), 3))
		// println(tempNum, sum)
		tempNum = tempNum / 10
	}

	if sum == num {
		check = true
		return
	}
	return
}
