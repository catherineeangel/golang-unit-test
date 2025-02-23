package logic02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal12(t *testing.T) {
	expected := [][]int{
		{1, 0, 5},
		{1, 3, 5},
		{1, 0, 5},
	}
	n := 3
	result := Soal12(n)

	for i, row := range result {
		assert.Equal(t, expected[i], row, "Soal12 failed at row %d", i)
	}
}
