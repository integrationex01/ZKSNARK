package example

import (
	"fmt"
	"testing"
)

func TestProo(t *testing.T) {
	css := Compilecircuit()
	constraints := css.GetConstraints()
	fmt.Printf("the constraint is %v\n", constraints)
	err := ProofGroth16(css)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("证明成功\n")

}
