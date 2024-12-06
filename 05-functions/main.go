package main

import "fmt"

func main() {
	hello() //function call
	greetings("alfaarghya") //function call with parameter

	x, y := 10, 11 
	fmt.Println(x,"+",y,"=", add(x,y) ) //function call with return value
	fmt.Println("----------")
	fmt.Println(x,"*",y,"=", multiply(x,y) ) //function call with return value
	fmt.Println("----------")

	a := 117
	val, ans := isEven(a) //function call with multiple return value
	fmt.Println(a,"% 2 =", val ) 
	fmt.Println(a,">", ans )
	fmt.Println("----------")

}

// a simple function
func hello() {
	fmt.Println("Greetings & Hello World! to U")	
	fmt.Println("----------")
}

//function with parameters
func greetings(name string) {
	fmt.Println("Greetings & Hello World! to ", name)
	fmt.Println("----------")
}

// function with return value
func add(a int, b int) int {
	sum := a+b
	return sum
}

//named return function
func multiply(a int, b int) (mul int) {
	mul = a*b
	return
}

//multiple named return in a function
func isEven(a int) (result int, evenOdd string) {
	result = a%2
	if (result == 0) {
		evenOdd = "it's even"
		return
	} else {
		evenOdd = "it's odd"
		return	
	}
}

