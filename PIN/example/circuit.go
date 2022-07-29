package example

import (
	"fmt"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

// Create the struct of circuit
type Circuit struct {
	A frontend.Variable `gnark:"x"` // x  --> secret visibility
}

func (c Circuit) Define(api frontend.API) error {
	// test ToBinary
	aB := api.ToBinary(c.A, 10)
	length := len(aB)
	Binary := []frontend.Variable{0, 0, 1, 0, 0, 1, 1, 0, 0, 0}
	for i := 0; i < length; i++ {
		api.AssertIsEqual(Binary[i], aB[i])
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
	// constraints := r1cs.GetConstraints()
	// fmt.Printf("r1cs:%v\n", r1cs)
	// fmt.Printf("the constraints are :%v\n", constraints)
	return r1cs
}
