package operations

import "fmt"


func Assignment() {
	var x = 11;

	fmt.Println("x ->",x);
	fmt.Println("----------");
}

func AddAssignment() {
	var x = 11;

	x += 5

	fmt.Println("x ->",x);
	fmt.Println("----------");
}

func SubtractAssignment() {
	var x = 11;

	x -= 3

	fmt.Println("x ->",x);
	fmt.Println("----------");
}

func MultiplyAssignment() {
	var x = 11;

	x *= 3

	fmt.Println("x ->",x);
	fmt.Println("----------");
}

func DivideAssignment() {
	var x = 11;

	x /= 3

	fmt.Println("x ->",x);
	fmt.Println("----------");
}

func ModulusAssignment() {
	var x = 11;

	x %= 3

	fmt.Println("x ->",x);
	fmt.Println("----------");
}

func BitwiseANDAssignment() {
	var x = 11;

	x &= 5

	fmt.Println("x ->",x);
	fmt.Println("----------");
}

func BitwiseORAssignment() {
	var x = 11;

	x |= 3

	fmt.Println("x ->",x);
	fmt.Println("----------");
}

func BitwiseXORAssignment() {
	var x = 11;

	x ^= 3

	fmt.Println("x ->",x);
	fmt.Println("----------");
}

func BitwiseLeftShiftAssignment() {
	var x = 11;

	x <<= 3

	fmt.Println("x ->",x);
	fmt.Println("----------");
}

func BitwiseRightShiftAssignment() {
	var x = 11;

	x <<= 3

	fmt.Println("x ->",x);
	fmt.Println("----------");
}