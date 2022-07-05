package main

import (
	"math"
)

// Cardinality Calculation
func getCardinality(set []float64) int {
	return len(set)
	// size := 0
	// for ind, el := range set {
	// 	size++
	// }
	// return size
}

// Mean Calculation
func getMean(set []float64) float64 {
	var mean float64
	for _, el := range set {
		mean = mean + el
	}
	return mean
}

// Variance Calculation
func getVariance(set []float64) float64 {
	var varinc, mean, sumVari float64
	mean = getMean(set)
	for _, el := range set {
		temp := math.Pow((el - mean), 2)
		sumVari = sumVari + math.Abs(temp)
	}
	varinc = sumVari / (float64(getCardinality(set)) - 1)
	return varinc
}

// Standard Deviation Calculation
func getStandardDeviation(set []float64) float64 {
	return math.Sqrt(getVariance(set))
}
