package logic01
// Soal 9
func Soal09 (n int) []int {
	slice:= make([]int, n)

	for i:= range (slice){
		if (i < n/2){
			// gausa sama dengan karena dia 0 indexing 
			(slice)[i] = 3 * (i+1)
			} else{
			(slice)[i] = 3 * (n - i)

		}
	}
	return slice
}