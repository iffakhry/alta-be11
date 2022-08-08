package score

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKonversiNilai(t *testing.T) {
	t.Run("Test nilai 100", func(t *testing.T) {
		expected := "Nilai A"
		actual := KonversiNilai(100)
		assert.Equal(t, expected, actual, "output tidak sesuai")
	})

	t.Run("Test nilai 70", func(t *testing.T) {
		expected := "Nilai B+"
		actual := KonversiNilai(70)
		assert.Equal(t, expected, actual, "output tidak sesuai")
	})
}
