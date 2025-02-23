package logic02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal11(t *testing.T) {
	expected := [][]int{
		{1, 3, 5},
		{0, 3, 5},
		{0, 0, 5},
	}
	n := 3
	result := Soal11(n)

	for i, row := range result {
		assert.Equal(t, expected[i], row, "Soal11 failed at row %d", i)
	}
}
