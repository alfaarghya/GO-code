package main

import (
	"fmt"
	"go-code/03-operations/operations"
)


func main() {
	
	fmt.Println("---------- Basic Operations ---------")
	operations.Add()
	operations.Subtract()
	operations.Multiply()
	operations.Divide()
	operations.Modulus()
	
	fmt.Println("---------- Comparison Operations ---------")
	operations.EqualTo()
	operations.NotEqual()
	operations.GreaterThan()
	operations.LessThan()
	operations.GreaterThanEqualTo()
	operations.LessThanEqualTo()
	
	fmt.Println("---------- Logical Operations ---------")
	operations.LogicalAND()
	operations.LogicalOR()
	operations.LogicalNOT()
	
	fmt.Println("---------- Bitwise Operations ---------")
	operations.BitwiseAND()
	operations.BitwiseOR()
	operations.BitwiseXOR()
	operations.BitwiseRightShift()
	operations.BitwiseLeftShift()
	
	fmt.Println("---------- Assignment Operations ---------")
	operations.Assignment()
	operations.AddAssignment()
	operations.SubtractAssignment()
	operations.MultiplyAssignment()
	operations.DivideAssignment()
	operations.ModulusAssignment()
	operations.BitwiseANDAssignment()
	operations.BitwiseORAssignment()
	operations.BitwiseXORAssignment()
	operations.BitwiseRightShiftAssignment()
	operations.BitwiseLeftShiftAssignment()
}