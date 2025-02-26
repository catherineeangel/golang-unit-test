package logic04

import (
	"fmt"
	"github.com/catherineeangel/golang-unit-test.git/utils"
)

func Soal05(n int) [][]int {
	fmt.Println("Soal 05")
	rowSize := n * (n + 1) / 2
	colSize := n * n
	result := utils.Create2DSlice(rowSize, colSize)

	row := 0
	col := 0
	for i := 1; i <= n; i++ {
		GeneratePyramid(row, col, i, result)

		row += i
		col += 2*(i-1) + 1
	}

	return result
}
