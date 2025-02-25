package logic03

import (
	"fmt"
	slice "github.com/catherineeangel/go-print-slice"
)

func Soal10(n int) (result [][]int) {
	fmt.Println(`Soal10`)
	result = slice.CreateSlice(n)

	// Fill the upper half including the middle row
	mid := (n - 1) / 2
	for row := range mid + 1 {
		num := (mid * 2) + 1
		for col := range mid + 1 {
			if row+col >= mid && col-row <= mid {
				//biar kebolak balik

				//ini sebelah kiri aja
				result[row][col] = num
				result[n-1-row][col] = num // Mirror to lower half

				//mirror ke kanan
				result[row][n-1-col] = num
				result[n-1-row][n-1-col] = num

				num -= 2 // Increase the number (odd sequence)
			}
		}
	}

	return result
}
