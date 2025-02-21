package logic01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoal03(t *testing.T) {
	slice := Soal03(6)
	expected := []int{3, 6, 9, 12, 15, 18}
	assert.Equal(t, len(slice), 6)
	for i, v := range expected {
		assert.Equal(t, slice[i], v)
	}
	assert.Equal(t, slice, expected)	
}

