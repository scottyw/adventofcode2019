package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Equal(t, []int{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, execute([]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}, 0))
	assert.Equal(t, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, execute([]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, 0))
}
