package logic01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSoal11(t *testing.T) {
	t.Run("Even n", func(t *testing.T) {
		slice:= Soal11(6)
		expected := []int{0, 1, 0, 3, 0, 6}
		assert.Equal(t, len(slice), 6)
		for i, v := range expected {
			assert.Equal(t, slice[i], v)
		}
	})

	t.Run("Odd n", func(t *testing.T) {
		slice:= Soal11(5)
		expected := []int{0, 1, 0, 3, 0}
		assert.Equal(t, len(slice), 5)
		for i, v := range expected {
			assert.Equal(t, slice[i], v)
		}
	})
}