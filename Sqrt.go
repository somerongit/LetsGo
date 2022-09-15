package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z:=1.0
	if x<100 {
		z=x/16
	} else if x<50 {
		z=x/8
	} else if x<20 {
		z=x/4
	}

    for i:=0;i<10;i++ {
		// Sqrt Formula(https://go.dev/tour/flowcontrol/8)
        z-=(z*z-x)/(2*z)

        fmt.Println(i,z)
		if z*z==x {
			break
		}
    }
    return z
}
