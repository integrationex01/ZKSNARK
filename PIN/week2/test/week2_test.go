package test

import (
	"testing"

	c "ZKSNARK/PIN/week2/circuit"

	"github.com/consensys/gnark/test"
)

func TestWeek2(t *testing.T) {
	// assert object wrapping testing.T
	assert := test.NewAssert(t)

	// declare circuit
	var circuit c.Circuit

	assert.ProverFailed(&circuit, &c.Circuit{
		X:   1,
		Y:   2,
		Z:   3,
		OUT: 7,
	})

	assert.ProverSucceeded(&circuit, &c.Circuit{
		X:   1,
		Y:   2,
		Z:   3,
		OUT: 6,
	})

	assert.ProverFailed(&circuit, &c.Circuit{
		X:   245,
		Y:   2,
		Z:   3,
		OUT: 6,
	})

	assert.ProverSucceeded(&circuit, &c.Circuit{
		X:   77,
		Y:   2,
		Z:   3,
		OUT: 1,
	})

}
