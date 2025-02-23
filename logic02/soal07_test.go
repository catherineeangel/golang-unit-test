package logic02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal07(t *testing.T) {
	expected := [][]int{
		{1, 0, 0},
		{0, 3, 0},
		{0, 0, 5},
	}
	n := 3
	result := Soal07(n)

	for i, row := range result {
		assert.Equal(t, expected[i], row, "Soal07 failed at row %d", i)
	}
}
