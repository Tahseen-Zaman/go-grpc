package main

// Orchestration of our program

import (
	"fmt"

	"github.com/Tahseen-Zaman/go-grpc/internal/adapters/core/arithmetic"
	"github.com/Tahseen-Zaman/go-grpc/internal/ports"
)

func main() {
	// ports
	var core ports.ArithmeticPort
	core = arithmetic.NewArithmeticAdapter()
	fmt.Println("Hello, World!")
	result, err := core.Addition(1, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	result, err = core.Subtraction(1, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	result, err = core.Addition(3, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	result, err = core.Division(1, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	result, err = core.Addition(3, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
	result, err = core.Subtraction(3, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}
