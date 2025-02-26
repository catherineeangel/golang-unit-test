package logic03

import (
	"fmt"
	slice "github.com/catherineeangel/go-print-slice"
)

func Soal06(n int) (result [][]int) {
	fmt.Println(`Soal06`)
	result = slice.CreateSlice(n)
	padding := 0
	counter := 1
	mid := n / 2
	for row := range mid + 1 {
		if row%2 == 0 {
			for col := 0 + padding; col < n-padding; col++ {
				result[row][col] = counter
				otherRow := n - 1 - row
				result[otherRow][col] = counter
				counter = counter + 2
			}
		} else {
			for col := n - padding - 1; col >= padding; col-- {
				result[row][col] = counter

				otherRow := n - 1 - row
				result[otherRow][col] = counter

				counter = counter + 2
			}
		}

		padding++
	}
	return result
}
