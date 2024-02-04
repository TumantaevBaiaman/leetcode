package leetcode

import "math/big"

func multiply(num1 string, num2 string) string {
	n1 := new(big.Int)
	n1.SetString(num1, 10)

	n2 := new(big.Int)
	n2.SetString(num2, 10)

	result := new(big.Int).Mul(n1, n2)
	return result.String()
}
