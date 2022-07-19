package main

import (
	c "ZKSNARK/PIN/week2/circuit"
	p "ZKSNARK/PIN/week2/proof"
	"fmt"
)

func main() {
	// 定义并编译电路
	css := c.Compilecircuit()
	err := p.ProofGroth16(css)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("证明成功")

}
