package operations

import "fmt"

func BitwiseAND() {
	x := 7
	y := 2

	fmt.Printf("x = 7 -> %b\n",x)
	fmt.Printf("y = 2 -> %b\n",y)

	fmt.Printf("x & y = %v -> ", x & y)
	fmt.Printf("%b\n" ,x & y)
	println("----------")
}

func BitwiseOR() {
	x := 7
	y := 2
	
	fmt.Printf("x = 7 -> %b\n",x)
	fmt.Printf("y = 2 -> %b\n",y)
	
	fmt.Printf("x | y = %v -> ", x | y)
	fmt.Printf("%b\n" ,x | y)
	println("----------")
}

func BitwiseXOR() {
	x := 7
	y := 2
	
	fmt.Printf("x = 7 -> %b\n",x)
	fmt.Printf("y = 2 -> %b\n",y)
	
	fmt.Printf("x ^ y = %v -> ", x ^ y)
	fmt.Printf("%b\n" ,x ^ y)
	println("----------")
}

func BitwiseRightShift() {
	x := 7
	y := 2
	
	fmt.Printf("x = 7 -> %b\n",x)
	fmt.Printf("y = 2 -> %b\n",y)
	
	fmt.Printf("x >> y = %v -> ", x >> y)
	fmt.Printf("%b\n" ,x << y)
	println("----------")
}

func BitwiseLeftShift() {
	x := 7
	y := 2
	
	fmt.Printf("x = 7 -> %b\n",x)
	fmt.Printf("y = 2 -> %b\n",y)
	
	fmt.Printf("x << y = %v -> ", x << y)
	fmt.Printf("%b\n" ,x >> y)
	println("----------")
}