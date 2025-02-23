package logic02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal06(t *testing.T) {
	expected := [][]int{
		{1, 4, 7},
		{14, 12, 10},
		{16, 19, 22},
	}
	n := 3
	result := Soal06(n)

	for i, row := range result {
		assert.Equal(t, expected[i], row, "Soal06 failed at row %d", i)
	}
}
