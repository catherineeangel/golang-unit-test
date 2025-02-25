package logic03

import (
	"fmt"
	slice "github.com/catherineeangel/go-print-slice"
)

func Soal11(n int) (result [][]int) {
	fmt.Println(`Soal11`)
	result = slice.CreateSlice(n)

	mid := (n - 1) / 2
	num := 1
	var reverseCol int
	for row := range mid + 1 {
		for col := range (result)[row] {

			if row == col {
				numLeft := (row * 2) + 1
				result[row][col] = numLeft
				result[n-row-1][col] = numLeft //mirror ke bwh
			}

			if row <= col {
				if row+col >= n-1 {
					if row%2 == 1 {
						result[row][col] = num
						result[n-row-1][col] = num // mirror ke bwh
					} else {
						reverseCol = n - 1 - (col - (n - 1 - row))
						result[row][reverseCol] = num
						result[n-row-1][reverseCol] = num // mirror ke bwh
					}
					num += 2
				}
			}

		}
	}
	return result
}
