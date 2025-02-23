package logic02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal10(t *testing.T) {
	expected := [][]int{
		{1, 0, 0},
		{1, 3, 0},
		{1, 3, 5},
	}
	n := 3
	result := Soal10(n)

	for i, row := range result {
		assert.Equal(t, expected[i], row, "Soal10 failed at row %d", i)
	}
}
