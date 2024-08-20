package main

import (
	"fmt"
	"math/big"
)

/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
значение которых > 2^20 (1048576).
*/

func main() {
	a := big.NewInt(54321054)
	b := new(big.Int)
	b.SetString("3215432", 10)

	fmt.Println(div(a, b).String())
	fmt.Println(mult(a, b).String())
	fmt.Println(sub(a, b).String())
	fmt.Println(sum(a, b).String())
}

func div(a, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Div(a, b)
	return result
}

func mult(a, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Mul(a, b)
	return result
}

func sub(a, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Sub(a, b)
	return result
}
func sum(a, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Add(a, b)
	return result
}
