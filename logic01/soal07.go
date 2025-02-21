package logic01

// Soal 7
func Soal07 (n int) []int {
	slice := make([]int, n)
	for i:= range (slice){
		if (i < n/2){
			// gausa sama dengan karena dia 0 indexing 
			(slice)[i] = 2 * i + 1
		} else{
			(slice)[i] = 2 * (n - i) - 1

		}
	}
	return slice
}
