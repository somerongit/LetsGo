package main

// GCD calculation

/*
GCD is learge number devide by number
iterate with range of small num
if num%count==0 && num1%count==0
then assign the count to gcd
*/

func getHcfBruteForce(num, num1 int64) int64 {
	if num == num1 {
		return num
	}

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

/*
GCD of 2 numbers remains the same,
not only by replacing the larger
number with the difference of those
numbers but also by replacing the
larger number with their sum.
*/
func getHcfNew(num, num1 int64) int64 {
	if num == num1 {
		return num
	}

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

// Euclid algo: gcd(a,b) => gcd(b,a%b)

func getHcfRecursion(num, num1 int64) int64 {
	if num == num1 {
		return num
	}

	if num1 == 0 {
		return num
	}
	return getHcfRecursion(num1, num%num1)
}

// LCM calculation

func getLcm(num, num1 int64) int64 {
	if num == num1 {
		return num
	}
	if num == 0 {
		return num1
	} else if num1 == 0 {
		return num
	}

	return (num * num1) / getHcfRecursion(num, num1)
}
