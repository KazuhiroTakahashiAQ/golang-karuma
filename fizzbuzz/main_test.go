package main_test

import (
	"testing"

	m "example.com/fizzbuzz"
	"github.com/stretchr/testify/assert"
)

func TestHoge(t *testing.T) {
	fizz := m.FizzBuzz(3)

	assert.Equal(t, "Fizz", fizz, "they should be equal")
}

func TestBuzz(t *testing.T) {
	buzz := m.FizzBuzz(5)

	assert.Equal(t, "Buzz", buzz, "ok?")
}

func TestFizzBuzz(t *testing.T) {
	fizzbuzz := m.FizzBuzz(15)
	assert.Equal(t, "FizzBuzz", fizzbuzz, "ok???")
}

func TestNone(t *testing.T) {
	none := m.FizzBuzz(49)
	assert.Equal(t, "49", none, "none?")
}
