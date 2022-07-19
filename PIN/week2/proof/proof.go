package proof

import (
	"fmt"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"

	c "ZKSNARK/PIN/week2/circuit"
)

func ProofGroth16(ccs frontend.CompiledConstraintSystem) error {
	// groth16 zkSNARK: Setup
	// generate prover key and vertifier key
	pk, vk, err := groth16.Setup(ccs)
	if err != nil {
		fmt.Println(err.Error())
	}

	// witness definition
	assignment := c.Circuit{X: 88, Y: 2, Z: 3, OUT: 1}
	witness, err := frontend.NewWitness(&assignment, ecc.BN254)
	if err != nil {
		fmt.Println(err.Error())
	}
	publicWitness, _ := witness.Public()

	// groth16: Prove & Verify
	// create proof by pk
	proof, err := groth16.Prove(ccs, pk, witness)
	if err != nil {
		fmt.Println(err.Error())
	}
	// vertifier proof by vk
	err = groth16.Verify(proof, vk, publicWitness)
	return err
}
