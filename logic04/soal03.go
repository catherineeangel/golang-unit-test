package logic04

import (
	"fmt"
	slice "github.com/catherineeangel/go-print-slice"
)

func Soal03(n int) [][]int {
	fmt.Println("Soal 03")
	arraySize := n * n
	result := slice.CreateSlice(arraySize)

	row := 0
	col := 0
	for i := 1; i <= n; i++ {
		GenerateBlock(row, col, i, result)
		if i != n {
			//	mirror ke kanan bawah and serong
			//	kiri bawah
			rowKiriBawah := arraySize - col - i
			GenerateBlock(rowKiriBawah, col, i, result)
			//	kanan atas
			colKananAtas := arraySize - row - i
			GenerateBlock(row, colKananAtas, i, result)
			//	kanan bawah
			GenerateBlock(rowKiriBawah, colKananAtas, i, result)
		}
		row += i
		col += i
	}

	return result
}
