package logic01

// Soal 8
func Soal08 (n int) []int {
	slice := make([]int, n)
	for i:= range (slice){
		if (i < n/2){
			(slice)[i] = 2 * (i+1)
			} else{
			(slice)[i] = 2 * (n - i)

		}
	}
	return slice
}
