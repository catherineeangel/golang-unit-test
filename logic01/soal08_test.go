package logic01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSoal08(t *testing.T) {
	t.Run("Even n", func(t *testing.T) {
		slice:= Soal08(6)
		expected := []int{2, 4, 6, 6, 4, 2}
		assert.Equal(t, len(slice), 6)
		for i, v := range expected {
			assert.Equal(t, slice[i], v)
		}
	})

	t.Run("Odd n", func(t *testing.T) {
		slice:= Soal08(5)
		expected := []int{2, 4, 6, 4, 2}
		assert.Equal(t, len(slice), 5)
		for i, v := range expected {
			assert.Equal(t, slice[i], v)
		}
	})
}