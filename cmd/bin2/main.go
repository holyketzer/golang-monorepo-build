package main

import (
	"example.com/gta_test/pkg/shared"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	fmt.Println(shared.Sum(11, 22))
	fmt.Println(shared.Mul(11, 22))
	fmt.Println(uuid.NewString())
}
