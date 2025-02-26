package logic04

import (
	"fmt"
	slice "github.com/catherineeangel/go-print-slice"
)

func Soal01(n int) [][]int {
	fmt.Println("Soal 01")
	arraySize := n * (n + 1) / 2
	result := slice.CreateSlice(arraySize)

	row := 0
	col := 0

	for i := 1; i <= n; i++ {
		GenerateBlock(row, col, i, result)
		row += i
		col += i
	}

	return result
}
