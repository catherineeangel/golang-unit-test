package logic02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal03(t *testing.T) {
	expected := [][]int{
		{1, 3, 5},
		{7, 9, 11},
		{13, 15, 17},
	}
	n := 3
	result := Soal03(n)

	for i, row := range result {
		assert.Equal(t, expected[i], row, "Soal03 failed at row %d", i)
	}
}
