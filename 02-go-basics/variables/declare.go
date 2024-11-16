package variables

import "fmt"


const PI float64 = 3.14 /*constant variable*/

func Declare () {
	var programming string = "GO"; /*declare with type*/
	var language = "go-lang"; /*declare without a type*/
	x := 1.618; /* can declare only with in a function */

	// programming = "hello" /*can change into other string*/
	// language = "world" /*can change into other string*/
	// x = 10 /*can change into other numeric values*/
	
	// programming = 101 /*can't change type after declare*/
	// language = 101 /*can't change type after declare*/
	// x = true /*can't change type after declare*/
	// PI = 10.10 /*can't change value or type after declare*/

	fmt.Println(">",programming)
	fmt.Println(">",language)
	fmt.Println(">",x)
	fmt.Println(">",PI)
	fmt.Println("----------")


}