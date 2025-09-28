package main

import(
	"fmt"
	"math/big"
)

func main() {
	a, b := new(big.Int), new(big.Int)

	a.SetString("12345678901234567890213213123213123123123123123123213123", 10)
	b.SetString("12345678901234567890213213123213123123123123123123213123", 10)

	fmt.Println(calc(a, b, '+'))
	fmt.Println(calc(a, b, '/'))
	fmt.Println(calc(a, b, '*'))
	fmt.Println(calc(a, b, '-'))
}

func calc(a, b *big.Int, op rune) *big.Int {
	switch op {
	case '+':
		return new(big.Int).Add(a, b)
	case '-':
		return new(big.Int).Sub(a, b)
	case '*':
		return new(big.Int).Mul(a, b)
	case '/':
		return new(big.Int).Div(a, b)
	default:
		return new(big.Int)
	}
}
