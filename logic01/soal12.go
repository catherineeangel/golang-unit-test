package logic01

// Soal 12
func Soal12(n int) []int  {
	slice:= make([]int, n)
	for i:= range (slice) {
		(slice)[i] = 2 * (i % 4) + 1
	}
	return slice
}