package calculator_test

import (
	"server/calculator"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	// Arrange
	var (
		num1, num2 int = 1, 2
		expected   int = 4
		calculator     = calculator.NewCalculator()
	)

	// Act
	actual := calculator.Sum(num1, num2)

	// Assert
	assert.Equal(t, actual, expected)
}
