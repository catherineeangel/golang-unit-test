package logic01

import "fmt"

// Soal 11
func Soal11(n int) []int {
	slice:= make([]int, n)
	initVal := 3
	for i := range n {
		if (i % 2 == 0){
			fmt.Print("Buzz", "\t")
		} else {
			switch i {
			case 1:
				fmt.Print(1, "\t")
				slice[i] = 1
			default: 
				fmt.Print(initVal, "\t")
				slice[i]=initVal
				initVal += 3

			}
		}
	}	
	return slice
}