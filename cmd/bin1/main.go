package main

import (
	"example.com/gta_test/pkg/shared"
	"fmt"
	"github.com/google/uuid"
)

func main() {
	fmt.Println(shared.Sum(1, 2))
	fmt.Println(uuid.NewString())
}
