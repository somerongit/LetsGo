package main

/*
Amicable numbers are found in pairs. A given pair of numbers
is Amicable if the sum of the proper divisors (not including itself)
of one number is equal to the other number and vice - versa.
For example 220 & 284 are amicable numbers
First we find the proper divisors of 220:
220 => 1, 2, 4, 5, 10, 11, 20, 22, 44, 55, 110
1+ 2 + 4 + 5 + 10 + 11 + 20 + 22 + 44 + 55 + 110 = 284
Now, 284 => 1, 2, 4, 71, 142
1 + 2 + 4 + 71 + 142 = 220

Other examples:
(1184, 1210), (2620, 2924), (5020, 5564), (6232, 6368),(10744, 10856),
(12285, 14595), (17296, 18416), and (63020, 76084)
*/

func checkAmicable(num, num1 int64) bool {
	sum := int64(0)
	for i := int64(1); i <= num/2; i++ {
		if num%i == 0 {
			sum = sum + i
		}
	}

	if sum == num1 {
		sum = 0
		for i := int64(1); i <= num1/2; i++ {
			if num1%i == 0 {
				sum = sum + i
			}
		}
		if sum == num {
			return true
		}
	}

	return false
}
