package operations

import "fmt"


func EqualTo() {
	x := 11;
	y := 111;
	// y := 11.11; /*invalid operation: x == y (mismatched types int and float64)*/

	str1 := "abc"
	str2 := "abc"
	// str2 := "xyz"

	fmt.Println("x == y ->",x == y);
	fmt.Println("str1 == str2 ->",str1 == str2);
	fmt.Println("----------");
}

func NotEqual() {
	x := 11;
	y := 110;
	// y := 11.11; /*invalid operation: x != y (mismatched types int and float64)*/

	str1 := "abc"
	str2 := "abc"
	// str2 := "xyz"
	
	fmt.Println("x != y ->",x != y);
	fmt.Println("str1 != str2 ->",str1 != str2);
	fmt.Println("----------");
}

func GreaterThan() {
	x := 11;
	y := 110;
	// y := 11.11; /*invalid operation: x > y (mismatched types int and float64)*/
	
	str1 := "abc"
	// str2 := "abc"
	str2 := "xyz"
	
	fmt.Println("x > y ->",x > y);
	fmt.Println("str1 > str2 ->",str1 > str2);
	fmt.Println("----------");
}

func LessThan() {
	x := 11;
	y := 110;
	// y := 11.11; /*invalid operation: x < y (mismatched types int and float64)*/
	
	str1 := "abc"
	// str2 := "abc"
	str2 := "xyz"
	
	fmt.Println("x < y ->",x < y);
	fmt.Println("str1 < str2 ->",str1 < str2);
	fmt.Println("----------");
}


func GreaterThanEqualTo() {
	x := 11;
	y := 110;
	// y := 11.11; /*invalid operation: x >= y (mismatched types int and float64)*/
	
	str1 := "abc"
	str2 := "abc"
	// str2 := "xyz"
	
	fmt.Println("x >= y ->",x >= y);
	fmt.Println("str1 >= str2 ->",str1 >= str2);
	fmt.Println("----------");
}

func LessThanEqualTo() {
	x := 11;
	y := 110;
	// y := 11.11; /*invalid operation: x <= y (mismatched types int and float64)*/
	
	str1 := "abc"
	str2 := "abc"
	// str2 := "xyz"
	
	fmt.Println("x <= y ->",x <= y);
	fmt.Println("str1 <= str2 ->",str1 <= str2);
	fmt.Println("----------");
}