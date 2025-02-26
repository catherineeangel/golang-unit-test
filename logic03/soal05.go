package logic03

import (
	"fmt"
	slice "github.com/catherineeangel/go-print-slice"
)

func Soal05(n int) (result [][]int) {
	fmt.Println(`Soal05`)
	result = slice.CreateSlice(n)
	start := 1
	mid := n / 2
	for row := range mid + 1 {
		for col := range mid + 1 {
			if row >= col {
				result[row][col] = start
				result[n-1-row][col] = start
				result[row][n-1-col] = start
				result[n-1-row][n-1-col] = start

				start += 2
			}
		}
	}
	return result
}
