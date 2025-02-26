package logic04

func GeneratePyramid(row int, col int, n int, matrix [][]int) {
	counter := 1
	left := n - 1
	jumlahCol := 2*(n-1) + 1

	for i := 0; i < n; i++ {
		if i%n == 0 {
			for j := jumlahCol - left - 1; j >= left; j-- {
				matrix[row+i][col+j] = counter
				counter += 2
			}

		} else {
			for j := left; j < jumlahCol-left; j++ {
				matrix[row+i][col+j] = counter
				counter += 2
			}
		}
		left--

	}

}

func GeneratePyramidReverse(row int, col int, n int, matrix [][]int) {
	counter := 1
	left := 0
	jumlahCol := 2*(n-1) + 1

	for i := 0; i < n; i++ {
		if i%n == 1 {
			for j := jumlahCol - left - 1; j >= left; j-- {
				matrix[row+i][col+j] = counter
				counter += 2
			}
		} else {
			for j := left; j < jumlahCol-left; j++ {
				matrix[row+i][col+j] = counter
				counter += 2
			}
		}
		left++

	}

}
