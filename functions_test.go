package gofp

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Map(t *testing.T) {
	t.Run("square int", func(t *testing.T) {
		o := Map([]int{0, 1, 2, 3}, func(i int) int { return i * i })

		assert.Equal(t, []int{0, 1, 4, 9}, o)
	})

	t.Run("double int", func(t *testing.T) {
		o := Map([]int{0, 1, 2, 3}, func(i int) int { return i * 2 })

		assert.Equal(t, []int{0, 2, 4, 6}, o)
	})

	t.Run("upper case string", func(t *testing.T) {
		o := Map([]string{"hello", "world"}, func(i string) string { return strings.ToUpper(i) })

		assert.Equal(t, []string{"HELLO", "WORLD"}, o)
	})
}

func Test_Filter(t *testing.T) {
	input := []int{-3, -2, -1, 0, 1, 2, 3}

	t.Run("numbers >= 0", func(t *testing.T) {
		o := Filter(input, func(i int) bool { return i >= 0 })

		assert.Equal(t, []int{0, 1, 2, 3}, o)
	})

	t.Run("numbers < 0", func(t *testing.T) {
		o := Filter(input, func(i int) bool { return i < 0 })

		assert.Equal(t, []int{-3, -2, -1}, o)
	})
}
