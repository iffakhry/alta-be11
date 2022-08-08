package calculator

/*
go test ./calculator/...

go test ./calculator/... -cover

go test ./calculator/... -coverprofile=cover.out && go tool cover -html=cover.out

=======
jika html tidak muncul otomatis:
go test ./calculator/... -coverprofile=cover.out && go tool cover -html=cover.out -o cover.html

=======
download testify:

go get -u github.com/stretchr/testify/assert
*/
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTambah(t *testing.T) {
	//test case
	t.Run("Test Fungsi Tambah bilangan positif", func(t *testing.T) {
		a := 10
		b := 20
		expected := 30
		actual := Tambah(a, b)
		if actual != expected {
			t.Error("hasil tambah tidak sesuai", actual)
		}
	})

	t.Run("Test Fungsi Tambah 2 bilangan negatif", func(t *testing.T) {
		a := -10
		b := -20
		expected := -30
		actual := Tambah(a, b)
		if actual != expected {
			t.Error("hasil tambah tidak sesuai", actual)
		}
	})

	t.Run("Test Fungsi Tambah bilangan ke1 negatif", func(t *testing.T) {
		a := -10
		b := 20
		expected := 10
		actual := Tambah(a, b)
		if actual != expected {
			t.Error("hasil tambah tidak sesuai", actual)
		}
	})

	t.Run("Test Fungsi Tambah bilangan ke2 negatif", func(t *testing.T) {
		a := 10
		b := -20
		expected := -10
		actual := Tambah(a, b)
		if actual != expected {
			t.Error("hasil tambah tidak sesuai", actual)
		}
	})
}

func TestKurang(t *testing.T) {
	t.Run("Test Fungsi Kurang bilangan positif", func(t *testing.T) {
		a := 20
		b := 10
		expected := 10
		actual := Kurang(a, b)
		assert.Equal(t, expected, actual, "hasil kurang tidak sesuai")
	})
	t.Run("Test Fungsi Kurang 2 bilangan negatif", func(t *testing.T) {
		expected := -10
		actual := Kurang(-20, -10)
		assert.Equal(t, expected, actual, "hasil kurang tidak sesuai")
	})
}
