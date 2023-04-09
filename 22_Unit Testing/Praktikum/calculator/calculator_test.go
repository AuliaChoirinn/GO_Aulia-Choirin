package calculator_test

import (
    "errors"
    "testing"

    "github.com/stretchr/testify/assert"

    "path/to/calculator"
)

func TestAddition(t *testing.T) {
    result := calculator.Addition(2, 3)
    assert.Equal(t, 5.0, result)
}

func TestSubtraction(t *testing.T) {
    result := calculator.Subtraction(5, 3)
    assert.Equal(t, 2.0, result)
}

func TestMultiplication(t *testing.T) {
    result := calculator.Multiplication(2, 3)
    assert.Equal(t, 6.0, result)
}

func TestDivision(t *testing.T) {
    t.Run("normal case", func(t *testing.T) {
        result, err := calculator.Division(6, 2)
        assert.Nil(t, err)
        assert.Equal(t, 3.0, result)
    })

    t.Run("division by zero", func(t *testing.T) {
        result, err := calculator.Division(6, 0)
        assert.Equal(t, float64(0), result)
        assert.Equal(t, errors.New("division by zero"), err)
    })
}
