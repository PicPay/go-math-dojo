package operations_test

import (
	"learningmath/pkg/operations"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	t.Run("should return 2", func(t *testing.T) {
		solution := operations.Add(1, 1)
		assert.Equal(t, 2, solution)
	})

	t.Run("should return 10", func(t *testing.T) {
		solution := operations.Add(5, 5)
		assert.Equal(t, 10, solution)

	})
}

func TestCustomAdd(t *testing.T) {
	t.Parallel()
	valorEsperado := 10
	valor := operations.Add(5, 5)
	if valor != valorEsperado {
		t.Error("O valor esperado é:", valorEsperado, "e o valor apresentado foi: ", valor)
	}

}
