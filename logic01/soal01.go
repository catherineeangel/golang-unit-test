package logic01

// Soal 1
func OddPattern(slice *[]int) {
	for i := range *slice {
		(*slice)[i] = 2*i + 1
	}
}

func Soal01(n int) []int {
	slice := make([]int, n)

	for i := range slice {
		(slice)[i] = 2*i + 1
	}
	return slice
}
