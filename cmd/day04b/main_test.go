package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	// 112233 meets these criteria because the digits never decrease and all repeated digits are exactly two digits long.
	// 123444 no longer meets the criteria (the repeated 44 is part of a larger group of 444).
	// 111122 meets the criteria (even though 1 is repeated more than twice, it still contains a double 22).
	assert.Equal(t, false, valid(111111))
	assert.Equal(t, true, valid(112233))
	assert.Equal(t, false, valid(123444))
	assert.Equal(t, true, valid(111122))
}
