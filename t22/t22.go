package main

import (
	"fmt"
	"math/big"
)

func main() {
	first := big.NewInt(10000000)
	second := big.NewInt(10000005)
	addResult := new(big.Int)
	addResult.Add(first, second)
	divResult := new(big.Int)
	divResult.Div(first, second)
	mulResult := new(big.Int)
	mulResult.Mul(first, second)
	subResult := new(big.Int)
	subResult.Sub(first, second)
	fmt.Println(mulResult, divResult, addResult, subResult)
}
