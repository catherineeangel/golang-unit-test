package logic01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoal01(t *testing.T) {
	slice := Soal01(5)

	assert.Equal(t, len(slice), 5)
	assert.Equal(t, slice[0], 1)
	assert.Equal(t, slice[1], 3)
}
