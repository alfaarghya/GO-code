package operations

import "fmt"

func LogicalAND() {
	x := 11;

	fmt.Println("x<10 && x>12 ->", x<10 && x>12);
	fmt.Println("----------");
}

func LogicalOR() {
	x := 11;

	fmt.Println("x<20 || x<10 ->", x<20 || x<10);
	fmt.Println("----------");
}

func LogicalNOT() {
	x := 11;

	fmt.Println("!(x<20 || x<10) ->", !(x<20 || x<10));
	fmt.Println("----------");
}