package main

/*
*********
_********
__*******
___******
____*****
_____****
______***
_______**
________*
*/

func getPattern(size int) {

	for j := 0; j < size; j++ {
		for k := 0; k < j; k++ {
			print("_")

		}

		for i := 0; i < size-j; i++ {
			print("*")
		}
		print("\n")
	}
}
