package xrange

import (
	s "slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXrange(t *testing.T) {
	result := s.Collect(Xrange(0, 5, 1))
	assert.Equal(t, []int{0, 1, 2, 3, 4}, result)
}

func TestXrangeBack(t *testing.T) {
	result := s.Collect(Xrange(5, 0, -1))
	assert.Equal(t, []int{5, 4, 3, 2, 1}, result)
}

func TestXrangeBackNegative(t *testing.T) {
	result := s.Collect(Xrange(-5, -10, -1))
	assert.Equal(t, []int{-5, -6, -7, -8, -9}, result)
}
func TestXrangeStep(t *testing.T) {
	result := s.Collect(Xrange(0, 10, 2))
	assert.Equal(t, []int{0, 2, 4, 6, 8}, result)
}

func TestXrangeNeg(t *testing.T) {
	result := s.Collect(Xrange(-6, -2, 2))
	assert.Equal(t, []int{-6, -4}, result)
}

func TestXrangeNegEmpty(t *testing.T) {
	result := s.Collect(Xrange(-2, -6, 2))
	assert.Equal(t, []int(nil), result)
}
func TestXrangeStepZero(t *testing.T) {
	result := s.Collect(Xrange(1, 23, 0))
	assert.Equal(t, []int(nil), result)
}

func TestXrangeEmpty(t *testing.T) {
	result := s.Collect(Xrange(0, 5, -1))
	assert.Equal(t, []int(nil), result)
}

func TestXrangeEmpty2(t *testing.T) {
	result := s.Collect(Xrange(5, 0, 1))
	assert.Equal(t, []int(nil), result)
}
