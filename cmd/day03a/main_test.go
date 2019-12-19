package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, 6, findClosestIntersection("R8,U5,L5,D3", "U7,R6,D4,L4"))
}
