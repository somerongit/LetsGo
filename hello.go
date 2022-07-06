package main

import (
	"fmt"
	"time"
)

func main() {
	var num int64
	print("Enter the range: ")
	fmt.Scan(&num)

	println("Input value is: ", num)

	print("\nSum of series for input range=> ")
	start := time.Now()
	println(getSeriesSum(num))
	end := time.Since(start)
	println("Time taken(BruteForce): ", end)

	// print("\nGCD of input values=> ")
	// start1 := time.Now()
	// println(getHcfNew(num))
	// end1 := time.Since(start1)
	// println("Time taken(New): ", end1)

	// print("\nGCD of input values=> ")
	// start2 := time.Now()
	// println(getHcfRecursion(num))
	// end2 := time.Since(start2)
	// println("Time taken(Recursion): ", end2)

}
