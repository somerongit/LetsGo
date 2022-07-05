package main

import (
	"fmt"
)

func main() {
	var size int
	print("Enter input size: ")
	fmt.Scan(&size)

	set := make([]float64, size)

	println("Enter the values: ")
	for i := 0; i < size; i++ {
		var temp float64
		fmt.Scan(&temp)
		set[i] = temp
	}
	println("Input values are")
	for i := 0; i < size; i++ {
		println(i, "-->", int(set[i]))
	}

	print("Cardinality of input value=> ")
	println(getCardinality(set))

	print("Mean of input value=> ")
	println(int64(getMean(set)))

	print("Variance of input value=> ")
	println(int64(getVariance(set)))

	print("Standard Deviation of input value=> ")
	println(int64(getStandardDeviation(set)))
}
