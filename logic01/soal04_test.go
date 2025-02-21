package logic01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoal04(t *testing.T) {
	slice := Soal04(6)
	expected := []int{11, 9, 7, 5, 3, 1}
	assert.Equal(t, len(slice), 6)
	for i, v := range expected {
		assert.Equal(t, slice[i], v)
	}
	assert.Equal(t, slice, expected)	
}

