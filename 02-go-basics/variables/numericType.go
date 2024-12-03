package variables

import "fmt"

func Uint8() {
	var value uint8 = 243 /* uint8 can contain values 0-255 */
	// var value uint8 = -1 /* cannot use -1 (untyped int constant) as uint8 value in variable declaration (overflows)*/
	// var value uint8 = 260 /*cannot use 260 (untyped int constant) as uint8 value in variable declaration (overflows)*/

	fmt.Println(">", value)
	fmt.Printf(">> Type is %T \n", value)
	fmt.Println(">>> unsigned  8-bit integers (0 to 255)")
	fmt.Println("----------")
}

func Uint16() {
	var value uint16 = 24300 /* uint16 can contain values 0-65535 */
	// var value uint16 = -1 /* cannot use -1 (untyped int constant) as uint16 value in variable declaration (overflows) */
	// var value uint16 = 70000 /* cannot use 70000 (untyped int constant) as uint16 value in variable declaration (overflows) */
	
	fmt.Println(">", value)
	fmt.Printf(">> Type is %T \n", value)
	fmt.Println(">>> unsigned 16-bit integers (0 to 65535)")
	fmt.Println("----------")
}

func Uint32() {
	var value uint32 = 429496295 /* uint32 can contain values 0-4294967295 */
	// var value uint32 = -1 /* cannot use -1 (untyped int constant) as uint32 value in variable declaration (overflows) */
	// var value uint32 = 4394967295 /* cannot use 4394967295 (untyped int constant) as uint32 value in variable declaration (overflows) */
	
	fmt.Println(">", value)
	fmt.Printf(">> Type is %T \n", value)
	fmt.Println(">>> unsigned 32-bit integers (0 to 4294967295)")
	fmt.Println("----------")
}

func Uint64() {
	var value uint64 = 1844674403751615 /* uint64 can contain values 0-4294967295 */
	// var value uint64 = -1 /* cannot use -1 (untyped int constant) as uint64 value in variable declaration (overflows) */
	// var value uint64 = 19446744073709551615 /* cannot use 4394967295 (untyped int constant) as uint64 value in variable declaration (overflows) */
	
	fmt.Println(">", value)
	fmt.Printf(">> Type is %T \n", value)
	fmt.Println(">>> unsigned 64-bit integers (0 to 18446744073709551615)")
	fmt.Println("----------")
}



func Int8() {
	var value int8 = 125 /* int8 can contain values -125 to 127 */
	// var value int8 = -125 /* int8 can contain values -125 to 127 */
	// var value int8 = -130 /* cannot use -130 (untyped int constant) as int8 value in variable declaration (overflows)*/
	// var value int8 = 260 /*cannot use 260 (untyped int constant) as int8 value in variable declaration (overflows)*/

	fmt.Println(">", value)
	fmt.Printf(">> Type is %T \n", value)
	fmt.Println(">>> signed  8-bit integers (-128 to 127)")
	fmt.Println("----------")
}

func Int16() {
	var value int16 = 30000 /* int16 can contain values -32768 to 32767 */
	// var value int16 = -30000 /* int16 can contain values -32768 to 32767 */
	// var value int16 = -35760 /* cannot use -35760 (untyped int constant) as int16 value in variable declaration (overflows)*/
	// var value int16 = 35760 /*cannot use 35760 (untyped int constant) as int16 value in variable declaration (overflows)*/

	fmt.Println(">", value)
	fmt.Printf(">> Type is %T \n", value)
	fmt.Println(">>> signed 16-bit integers (-32768 to 32767)")
	fmt.Println("----------")
}

func Int32() {
	var value int32 = 2147483000 /* int32 can contain values -2147483648 to 2147483647 */
	// var value int32 = -2147483000 /* int32 can contain values -2147483648 to 2147483647 */
	// var value int32 = -2247483000 /* cannot use -2247483000 (untyped int constant) as int32 value in variable declaration (overflows)*/
	// var value int32 = 2247483000 /*cannot use 2247483000 (untyped int constant) as int32 value in variable declaration (overflows)*/

	fmt.Println(">", value)
	fmt.Printf(">> Type is %T \n", value)
	fmt.Println(">>> signed 32-bit integers (-2147483648 to 2147483647)")
	fmt.Println("----------")
}

func Int64() {
	var value int64 = 214748789000 /* int64 can contain values -9223372036854775808 to 9223372036854775807 */
	// var value int64 = -9223362036854775808 /* int64 can contain values -9223372036854775808 to 9223372036854775807 */
	// var value int64 = -9323372036854775808 /* cannot use -9323372036854775808 (untyped int constant) as int64 value in variable declaration (overflows)*/
	// var value int64 = 9323372036854775808 /*cannot use 9323372036854775808 (untyped int constant) as int64 value in variable declaration (overflows)*/

	fmt.Println(">", value)
	fmt.Printf(">> Type is %T \n", value)
	fmt.Println(">>> signed 64-bit integers (-9223372036854775808 to 9223372036854775807)")
	fmt.Println("----------")
}
	
	func Int() {
		var value int = 9200000000000000007 

		fmt.Println(">", value)
		fmt.Printf(">> Type is %T \n", value)
		fmt.Println(">>> It's general way declaring a Integer")
		fmt.Println("---------- ----------")
}


func Float32() {
	var value float32 = 2147.123456789
	
	fmt.Println(">", value)
	fmt.Printf(">> Type is %T \n", value)
	fmt.Println(">>> IEEE 754 32-bit floating-point numbers")
	fmt.Println("---------- ----------")
}

func Float64() {
	var value float64 = 2147.1234567890123456789
	
	fmt.Println(">", value)
	fmt.Printf(">> Type is %T \n", value)
	fmt.Println(">>> IEEE 754 64-bit floating-point numbers")
	fmt.Println("---------- ----------")
}