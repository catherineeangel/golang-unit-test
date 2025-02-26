package logic03

import (
	"fmt"
	slice "github.com/catherineeangel/go-print-slice"
)

func Soal14(n int) (result [][]int) {
	fmt.Println(`Soal14`)
	result = slice.CreateSlice(n)
	start := 1
	for row := range result {
		for col := range result[row] {
			if row%2 == 0 {
				result[col][row] = start
			} else {
				result[n-1-col][row] = start
			}
			start += 2
		}
	}

	return result
}
