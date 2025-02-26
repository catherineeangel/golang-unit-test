package logic03

import (
	"fmt"
	slice "github.com/catherineeangel/go-print-slice"
)

func Soal13(n int) (result [][]int) {
	fmt.Println(`Soal13`)
	result = slice.CreateSlice(n)

	// Fill the upper half including the middle row
	mid := n / 2
	//num := 1
	for row := range mid + 1 {
		for col := range mid + 1 {
			//if row+col >= mid && col-row <= mid {
			//	//biar kebolak balik
			//	if col%2 == 1 {
			//		result[col][row] = num
			//		result[col][n-1-row] = num // Mirror to lower half
			//	} else {
			//		result[n-1-col][row] = num
			//		result[n-1-col][n-1-row] = num
			//	}
			//	//num += 2 // Increase the number (odd sequence)
			//}
			if ((row+col)%2 == 0) && (row+col >= 4) {
				result[row][col] = 2*col + 1
				result[n-1-row][col] = 2*col + 1
				result[row][n-1-col] = 2*col + 1
				result[n-1-row][n-1-col] = 2*col + 1
			}
		}
	}

	return result
}
