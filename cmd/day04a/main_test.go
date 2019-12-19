package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	// 111111 meets these criteria (double 11, never decreases).
	// 223450 does not meet these criteria (decreasing pair of digits 50).
	// 123789 does not meet these criteria (no double).

	assert.Equal(t, true, valid(111111))
	assert.Equal(t, false, valid(223450))
	assert.Equal(t, false, valid(123789))
}
