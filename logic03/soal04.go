package logic03

import (
	"fmt"
	slice "github.com/catherineeangel/go-print-slice"
)

func Soal04(n int) (result [][]int) {
	fmt.Println(`Soal04`)
	result = slice.CreateSlice(n)
	start := 1
	var reverseCol int
	for row := range result {
		for col := range result[row] {
			if row+col >= n-1 {
				if row%2 == 0 {

					result[row][col] = start
				} else {
					reverseCol = n - 1 - (col - (n - 1 - row))
					result[row][reverseCol] = start
				}
				start += 2
			}

		}
	}
	return result
}
