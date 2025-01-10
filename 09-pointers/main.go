package main

import "fmt"

func main() {
	/*var x int = 10
	fmt.Println(x)    // output -> 10
	changeTheValue(x) // output -> 20
	fmt.Println(x)    // output -> 10*/

	var x int = 10
	var ptr = &x      //ptr is address reference of x
	fmt.Println(ptr)  // it's shows the memory address of x
	fmt.Println(*ptr) // it's shows the the value in the memory address -> 10
	fmt.Println("------------------")

	fmt.Println(x)      // output -> 10
	changeTheValue(ptr) // output -> 20
	fmt.Println(x)      // output -> 20
}

/*func changeTheValue(x int) {
	x = 20
	fmt.Println(x)
}*/

func changeTheValue(ptr *int) {
	*ptr = 20 // change the actual value of x
	fmt.Println(*ptr)
}
