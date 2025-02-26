package logic03

import (
	"fmt"
	slice "github.com/catherineeangel/go-print-slice"
)

func Soal03(n int) (result [][]int) {
	fmt.Println(`Soal03`)
	result = slice.CreateSlice(n)
	start := 1
	for i := range result {
		for j := range (result)[i] {
			if i+j <= n-1 {
				if i%2 == 0 {
					(result)[i][j] = start
					start += 2
				} else {
					// karena nilai max nya di n
					reverseCol := n - 1 - j - i
					(result)[i][reverseCol] = start
					start += 3
				}
			}
		}
	}
	return result
}
