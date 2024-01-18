package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {

	var a string = "Hello"
	var b string = "Hello"
	assert.Equal(t, a, b, "The two worlds will be the same.")

	assert.Equal(t, 123, 123, "they should be equal")
	assert.NotEqual(t, 124, 456, "should not be equal")

	assert.True(t, true, "True is true")

	assert2 := assert.New(t)
	assert2.Equal(123, 123, "they should be equal")

	assert2.Equal(a, b, "The two worlds will be the same.")

}
