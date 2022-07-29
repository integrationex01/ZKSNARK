package circuit

import (
	"fmt"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

// Create the struct of circuit
type Circuit struct {
	A frontend.Variable `gnark:"x"`   // x  --> secret visibility
	B frontend.Variable `gnark:"y"`   // y  --> secret visibility
	C frontend.Variable `gnark:"out"` // out --> public visibility
}

func (c Circuit) Define(api frontend.API) error {
	// String to Binary
	cB := api.ToBinary(c.C, 3)
	length := len(cB)

	aB := api.ToBinary(c.A, length)
	bB := api.ToBinary(c.B, length)

	var upbool frontend.Variable = 0

	// var resi frontend.Variable

	// Add
	for i := 0; i < length; i++ {
		resi := api.Xor(upbool, api.Xor(aB[i], bB[i]))
		api.AssertIsEqual(cB[i], resi)
		upbool = api.Or(api.And(aB[i], bB[i]), api.And(upbool, api.Xor(aB[i], bB[i])))
	}

	return nil
}

// flatten circuit to R1CS form
func Compilecircuit() frontend.CompiledConstraintSystem {
	var MyCircuit Circuit
	r1cs, err := frontend.Compile(ecc.BN254, r1cs.NewBuilder, &MyCircuit)
	if err != nil {
		fmt.Printf("error : %s\n", err.Error())
	}
	return r1cs
}
