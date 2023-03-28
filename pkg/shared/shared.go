package shared

import "fmt"

//go:noinline
func Sum(a, b int) int {
	res := a + b
	fmt.Println("Call sum:", res)
	return res
}

//go:noinline
func Mul(a, b int) int {
	res := a * b
	fmt.Println("Call mul:", res)
	return res
}
