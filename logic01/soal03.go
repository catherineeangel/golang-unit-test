package logic01

// Soal 3
func Soal03(n int) []int  {
	slice := make([]int, n)
	for i:= range slice {
		(slice)[i] = 3 * (i + 1)
	}
	return slice
}


