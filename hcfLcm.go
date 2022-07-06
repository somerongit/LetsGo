package main

// GCD calculation

/* gcd is learge no devide by number
iterate with range of small num
till num%count==0 && num1%count==0
count is gcd
*/

func getHcfBruteForce(num, num1 int64) int64 {
	if num == 0 {
		return num1
	} else if num1 == 0 {
		return num
	}

	var temp int64
	if num > num1 {
		temp = num1
	} else {
		temp = num
	}

	var i, count int64
	for i = 1; i <= temp; i++ {
		if num%i == 0 && num1%i == 0 {
			count = i
		}
	}

	return count
}

func getHcfNew(num, num1 int64) int64 {
	if num == 0 {
		return num1
	}

	for num1 != 0 {
		if num > num1 {
			num = num - num1
		} else {
			num1 = num1 - num
		}
	}

	return num
}

func getHcfRecursion(num, num1 int64) int64 {
	if num1 == 0 {
		return num
	}
	return getHcfRecursion(num1, num%num1)
}

// LCM calculation

func getLcm(num, num1 int64) int64 {
	return (num * num1) / getHcfRecursion(num, num1)
}
