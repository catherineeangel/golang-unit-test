package logic01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoal07(t *testing.T) {
	t.Run("Even n", func(t *testing.T) {
		slice:= Soal07(6)
		expected := []int{1, 3, 5, 5, 3, 1}
		assert.Equal(t, len(slice), 6)
		for i, v := range expected {
			assert.Equal(t, slice[i], v)
		}
	})

	t.Run("Odd n", func(t *testing.T) {
		slice:= Soal07(5)
		expected := []int{1, 3, 5, 3, 1}
		assert.Equal(t, len(slice), 5)
		for i, v := range expected {
			assert.Equal(t, slice[i], v)
		}
	})	
}

