package logic01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSoal12(t *testing.T) {
	t.Run("Even n", func(t *testing.T) {
		slice:= Soal12(4)
		expected := []int{1, 3, 5, 7}
		assert.Equal(t, len(slice), 4)
		for i, v := range expected {
			assert.Equal(t, slice[i], v)
		}
	})

	t.Run("Odd n", func(t *testing.T) {
		slice:= Soal12(5)
		expected := []int{1, 3, 5, 7, 1}
		assert.Equal(t, len(slice), 5)
		for i, v := range expected {
			assert.Equal(t, slice[i], v)
		}
	})
}