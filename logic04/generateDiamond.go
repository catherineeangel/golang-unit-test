package logic04

func GenerateDiamond(row int, col int, n int, matrix [][]int) {
	counter := 1
	left := n - 1
	jumlahCol := 2*(n-1) + 1
	jumlahRow := jumlahCol

	for i := 0; i < jumlahRow; i++ {
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
		if i < jumlahRow/2 {
			left--
		} else {
			left++
		}

	}

}
