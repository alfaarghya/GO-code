package operations

import "fmt"

func Add() {
	x := 11;
	y := 111;
	// y := 11.11; /* can't do this -> invalid operation: x + y (mismatched types int and float64)*/

	str1 := "abc"
	str2 := "xyz"
	// str2 := 123 /*invalid operation: str1 + str2 (mismatched types string and int)*/


	fmt.Println("x+y =",x+y);
	fmt.Println("str1+str2 =",str1+str2);
	fmt.Println("---------");
}

func Subtract() {
	x := 11;
	y := 111;
	// y := 11.11; /* can't do this -> invalid operation: x - y (mismatched types int and float64)*/

	// str1 := "abc"
	// str2 := "xyz"


	fmt.Println("x-y =",x-y);
	fmt.Println("---------");
	// fmt.Println("str1-str2 =",str1-str2); /*invalid operation: operator - not defined on str1 (variable of type string)*/
}

func Multiply() {
	x := 11;
	y := 111;
	// y := 11.11; /*invalid operation: x * y (mismatched types int and float64)*/

	fmt.Println("x*y =",x*y);
	fmt.Println("----------");
}

func Divide() {
	x := 11;
	y := 111;
	// y := 11.11; /*invalid operation: x / y (mismatched types int and float64)*/

	fmt.Println("x/y =",x/y);
	fmt.Println("----------");
}

func Modulus() {
	x := 11;
	y := 111;
	// y := 11.11; /*invalid operation: x % y (mismatched types int and float64)*/

	fmt.Println("x%y =",x%y);
	fmt.Println("----------");
}
