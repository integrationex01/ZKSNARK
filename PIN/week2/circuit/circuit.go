package circuit

import (
	"fmt"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

// Create the struct of circuit
type Circuit struct {
	X   frontend.Variable `gnark:"x"`   // x  --> secret visibility
	Y   frontend.Variable `gnark:"y"`   // y  --> secret visibility
	Z   frontend.Variable `gnark:"z"`   // z  --> secret visibility
	OUT frontend.Variable `gnark:"out"` // out --> public visibility
}

// Select if b is true, yields i1 else yields i2
// func (circuit *Circuit) Select(b frontend.Variable, i1, i2 interface{}) frontend.Variable {
// 	if b == 1

// }
// Define the form of circuit
func (circuit *Circuit) Define(api frontend.API) error {
	// compute c*x and store it in the local variable cx.
	// cx := api.Mul(circuit.C, circuit.X)

	// compute y*z and store it in the local variable yz.
	yz := api.Mul(circuit.Y, circuit.Z)

	// compute c*x*y*z and store it in the local variable cxyz.
	// cxyz := api.Mul(cx, yz)

	// compute 2*y-z and store it in the local variable yz.
	ypyrz := api.Sub(api.Mul(2, circuit.Y), circuit.Z)

	// compute 1-c*x and store it in the local variable yz.
	// onercx := api.Mul(-1, api.Add(cx, -1))

	// compute (1-c*x)*(2*y-z)+(c*x*y*z) and store it in the local variable res
	// res := api.Add(api.Mul(onercx, ypyrz), cxyz)
	// var c int
	// var res frontend.Variable

	// if circuit.X == 1 {
	// 	c = 1
	// 	res = api.Mul(c, api.Mul(circuit.X, yz))
	// } else {
	// 	c = 0
	// 	res = api.Mul(api.Mul(-1, api.Add(api.Mul(c, circuit.X), -1)), ypyrz)
	// }
	res := api.Select(api.IsZero(api.Sub(circuit.X, 1)), yz, ypyrz)
	// assert that the statement  is true.
	api.AssertIsEqual(circuit.OUT, res)
	return nil
}

// func btoi(b bool) int {
// 	if b {
// 		return 1
// 	}
// 	return 0
// }

// flatten circuit to R1CS form
func Compilecircuit() frontend.CompiledConstraintSystem {
	var MyCircuit Circuit
	r1cs, err := frontend.Compile(ecc.BN254, r1cs.NewBuilder, &MyCircuit)
	if err != nil {
		fmt.Printf("error : %s\n", err.Error())
	}
	fmt.Printf("rics:%v\n", r1cs)
	return r1cs
}
