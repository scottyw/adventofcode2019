package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// For a mass of 12, divide by 3 and round down to get 4, then subtract 2 to get 2.
// For a mass of 14, dividing by 3 and rounding down still yields 4, so the fuel required is also 2.
// For a mass of 1969, the fuel required is 654.
// For a mass of 100756, the fuel required is 33583.

func TestCalculateFuel(t *testing.T) {
	assert.Equal(t, 2, calculateFuel(12))
	assert.Equal(t, 2, calculateFuel(14))
	assert.Equal(t, 654, calculateFuel(1969))
	assert.Equal(t, 33583, calculateFuel(100756))
}
