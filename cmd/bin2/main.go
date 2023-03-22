package main

import (
	"fmt"
	"github.com/google/uuid"
)
import "example.com/gta_test/pkg/shared"

func main() {
	fmt.Println(shared.Sum(11, 22))
	fmt.Println(uuid.NewString())
	fmt.Println(uuid.NewString())
}
