package logic04

import (
	"fmt"
	slice "github.com/catherineeangel/go-print-slice"
)

func Soal08(n int) [][]int {
	fmt.Println("Soal 08")
	size := n * n
	result := slice.CreateSlice(size)

	row := 0
	col := 0
	for i := 1; i <= n; i++ {
		GenerateDiamond(row, col, i, result)
		row += 2*(i-1) + 1
		col += 2*(i-1) + 1
	}

	return result
}
