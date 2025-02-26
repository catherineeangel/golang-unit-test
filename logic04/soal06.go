package logic04

import (
	"fmt"
	"github.com/catherineeangel/golang-unit-test.git/utils"
)

func Soal06(n int) [][]int {
	fmt.Println("Soal 06")
	rowSize := n * (n + 1) / 2
	colSize := n * n
	result := utils.Create2DSlice(rowSize, colSize)

	row := rowSize - 1
	col := 0
	for i := 1; i <= n; i++ {
		GeneratePyramid(row, col, i, result)
		row -= i + 1
		col += 2*(i-1) + 1
	}

	return result
}
