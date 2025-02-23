package logic02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal02(t *testing.T) {
	expected := [][]int{
		{2, 4, 6},
		{2, 4, 6},
		{2, 4, 6},
	}
	n := 3
	result := Soal02(n)

	for i, row := range result {
		assert.Equal(t, expected[i], row, "Soal02 failed at row %d", i)
	}
}
