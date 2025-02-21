package logic01

// import "github.com/catherineeangel/golang-exercise/utils"
import (
	"fmt"

	slice "github.com/catherineeangel/go-print-slice"
)

func Sample01() {
	array := generateSlice(10)
	n := 10
	fmt.Println("Odd Slice")
	array = Soal01(2)
	slice.Print1DSlice(array)
	fmt.Println("Even Slice")
	Soal02(n)
	slice.Print1DSlice(array)
	fmt.Println("Modulo Four Pattern")
	Soal12(n)
	slice.Print1DSlice(array)
	fmt.Println("Reverse Odd Pattern")
	Soal04(n)
	slice.Print1DSlice(array)

	fmt.Println("Repeating Odd Pattern")
	Soal07(n)
	slice.Print1DSlice(array)
	fmt.Println("Repeating Even Pattern")
	Soal08(n)
	slice.Print1DSlice(array)
	fmt.Println("Repeating Triple Pattern")
	Soal09(n)
	slice.Print1DSlice(array)

	fmt.Println("Buzz Pattern")
	Soal11(7)

	fmt.Println()

}

func generateSlice(n int) []int {
	return make([]int, n)
}
