package main

import (
	"fmt"
)

func main() {
	var num int64
	fmt.Scanf("%d", &num)
	print(checkArmstrong(num))
}
