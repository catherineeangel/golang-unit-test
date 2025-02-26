package logic04

func GenerateBlock(row int, col int, n int, matrix [][]int) {
	counter := 1

	for i := row; i < row+n; i++ {
		for j := col; j < col+n; j++ {
			matrix[i][j] = counter
			counter += 2
		}
	}
}
