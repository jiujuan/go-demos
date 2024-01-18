package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	assert.Equal(t, 123, 123, "they should be equal")
	assert.NotEqual(t, 124, 456, "should not be equal")

	assert2 := assert.New(t)
	assert2.Equal(123, 123, "they should be equal")
}
