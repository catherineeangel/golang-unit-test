package logic01


func Soal01(n int) []int {
	slice := make([]int, n)

	for i := range slice {
		(slice)[i] = 2*i + 1
	}
	return slice
}
