package logic04

import (
	"fmt"
	slice "github.com/catherineeangel/go-print-slice"
)

func Soal02(n int) [][]int {
	fmt.Println("Soal 02")
	arraySize := n * (n + 1) / 2
	result := slice.CreateSlice(arraySize)

	row := arraySize - 1
	col := 0

	for i := 1; i <= n; i++ {
		GenerateBlock(row, col, i, result)
		row -= i + 1
		col += i
	}

	return result
}
