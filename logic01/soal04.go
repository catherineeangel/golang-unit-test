package logic01

// Soal 4
func Soal04 (n int) []int  {
	slice := make([]int, n)
	for i:= range (slice) {
		(slice)[i] = (2 * (n - i - 1) + 1)
	}
	return slice
}
