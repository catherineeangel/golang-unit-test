package logic02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal08(t *testing.T) {
	expected := [][]int{
		{0, 0, 5},
		{0, 3, 0},
		{1, 0, 0},
	}
	n := 3
	result := Soal08(n)

	for i, row := range result {
		assert.Equal(t, expected[i], row, "Soal08 failed at row %d", i)
	}
}
