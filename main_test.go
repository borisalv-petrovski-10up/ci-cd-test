package main_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	msg := "hello world"
	assert.Equal(t, msg, "hello world")
}
