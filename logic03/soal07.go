package logic03

import (
	"fmt"
	slice "github.com/catherineeangel/go-print-slice"
)

func Soal07(n int) (result [][]int) {
	fmt.Println(`Soal07`)
	result = slice.CreateSlice(n)

	// Fill the upper half including the middle row
	mid := (n - 1) / 2
	fmt.Println(mid)
	num := 1

	for row := range mid + 1 {
		for col := range result[row] {
			if row+col >= mid && col-row <= mid {
				//biar kebolak balik
				if row%2 == 1 {
					result[row][col] = num
					result[n-1-row][col] = num // Mirror to lower half
				} else {
					result[row][n-1-col] = num
					result[n-1-row][n-1-col] = num // mirror ke lower half
				}
				num += 2 // Increase the number (odd sequence)
			}
		}
	}

	return result
}
