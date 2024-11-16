package main

import (
	"fmt"
	"go-code/02-go-basics/variables"
)

func main() {

	fmt.Println("---------- Variable Declaration ---------")
	variables.Declare()
	
	fmt.Println("---------- Numeric Types ---------")
	variables.Uint8()
	variables.Uint16()
	variables.Uint32()
	variables.Uint64()
	
	variables.Int8()
	variables.Int16()
	variables.Int32()
	variables.Int64()
	
	variables.Int()
	variables.Float32()
	variables.Float64()

	fmt.Println("---------- Boolean Type ---------")
	variables.Boolean()
	
	fmt.Println("---------- String Type ---------")
	variables.String()
}