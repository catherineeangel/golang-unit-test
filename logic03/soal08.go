package logic03

import (
	"fmt"
	slice "github.com/catherineeangel/go-print-slice"
)

func Soal08(n int) (result [][]int) {
	fmt.Println(`Soal08`)
	result = slice.CreateSlice(n)

	// Fill the upper half including the middle row
	mid := (n - 1) / 2
	num := 1
	for row := range mid + 1 {
		for col := range result[row] {
			if row+col >= mid && col-row <= mid {
				//biar kebolak balik
				if col%2 == 1 {
					result[col][row] = num
					result[col][n-1-row] = num // Mirror to lower half
				} else {
					result[n-1-col][row] = num
					result[n-1-col][n-1-row] = num
				}
				num += 2 // Increase the number (odd sequence)
			}
		}
	}

	return result
}
