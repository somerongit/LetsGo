package main

/*
Sum of the series 1+11+111+....111 upto n
For eg. if n=4, the series is : 1+11+111+1111
*/

func getSeriesSum(num int64) int64 {
	if num == 0 {
		return int64(0)
	} else if num == 1 {
		return int64(1)
	}
	sum := int64(1)
	for i := int64(2); i <= num; i++ {
		sum = sum*10 + i
	}
	return sum
}
