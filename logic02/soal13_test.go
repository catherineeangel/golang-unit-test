package logic02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal13(t *testing.T) {
	expected := [][]int{
		{1, 3, 5},
		{0, 3, 0},
		{1, 3, 5},
	}
	n := 3
	result := Soal13(n)

	for i, row := range result {
		assert.Equal(t, expected[i], row, "Soal13 failed at row %d", i)
	}
}
