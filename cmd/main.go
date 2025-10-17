package main

// Orchestration of our program

import (
	"fmt"
	"log"
	"os"

	api "github.com/Tahseen-Zaman/go-grpc/internal/adapters/app/api"
	"github.com/Tahseen-Zaman/go-grpc/internal/adapters/core/arithmetic"
	gRPC "github.com/Tahseen-Zaman/go-grpc/internal/adapters/framework/left/grpc"
	"github.com/Tahseen-Zaman/go-grpc/internal/adapters/framework/right/db"
	"github.com/Tahseen-Zaman/go-grpc/internal/ports"
)

func main() {
	// ports
	var dbaseAdapter ports.DBPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort
	var err error

	dbaseDriver := os.Getenv("DB_DRIVER")
	dbaseSource := os.Getenv("DB_SOURCE")
	dbaseAdapter, err = db.NewDBAdapter(dbaseDriver, dbaseSource)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
		return
	}
	defer dbaseAdapter.CloseConnection()

	core = arithmetic.NewArithmeticAdapter()

	appAdapter = api.NewAPIAdapter(core, dbaseAdapter)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)

	gRPCAdapter.Run()

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
