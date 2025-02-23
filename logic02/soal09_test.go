package logic02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal09(t *testing.T) {
	expected := [][]int{
		{1, 0, 5},
		{0, 3, 0},
		{1, 0, 5},
	}
	n := 3
	result := Soal09(n)

	for i, row := range result {
		assert.Equal(t, expected[i], row, "Soal09 failed at row %d", i)
	}
}
