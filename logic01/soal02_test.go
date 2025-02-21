package logic01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoal02(t *testing.T) {
	slice := Soal02(5)
	expected := []int{2, 4, 6, 8, 10}
	assert.Equal(t, len(slice), 5)
	assert.Equal(t, slice[1], expected[1])
	assert.Equal(t, slice[0], expected[0])
	assert.Equal(t, slice, expected)
}