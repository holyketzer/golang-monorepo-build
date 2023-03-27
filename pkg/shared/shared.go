package shared

import "fmt"

func Sum(a, b int) int {
	res := a + b
	fmt.Println("Call sum:", res)
	return res
}
