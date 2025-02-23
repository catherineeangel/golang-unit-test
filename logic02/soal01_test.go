package logic02

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSoal01(t *testing.T) {

	expected := [][]int{
		{1, 3, 5},
		{1, 3, 5},
		{1, 3, 5},
	}
	n := 3
	result := Soal01(n)

	for i, row := range result {
		assert.Equal(t, expected[i], row, "Soal01 failed at row %d", i)
	}
}
