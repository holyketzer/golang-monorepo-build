package main

import (
	"fmt"
	"github.com/google/uuid"
	"example.com/gta_test/pkg/shared"
)

func main() {
	fmt.Println(shared.Sum(1, 2))
	fmt.Println(uuid.NewString())
}
