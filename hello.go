package main

import (
	"fmt"
	"time"
)

func main() {
	var num, num1 int64
	print("Enter 1st num: ")
	fmt.Scan(&num)

	print("Enter 2nd num: ")
	fmt.Scan(&num1)

	println("Input values are: ", num, ",", num1)

	print("\nGCD of input values=> ")
	start := time.Now()
	println(getHcfBruteForce(num, num1))
	end := time.Since(start)
	println("Time taken(BruteForce): ", end)

	print("\nGCD of input values=> ")
	start1 := time.Now()
	println(getHcfNew(num, num1))
	end1 := time.Since(start1)
	println("Time taken(New): ", end1)

	print("\nGCD of input values=> ")
	start2 := time.Now()
	println(getHcfRecursion(num, num1))
	end2 := time.Since(start2)
	println("Time taken(Recursion): ", end2)

	print("\nLCM of input values=> ")
	print(getLcm(num, num1))

}
