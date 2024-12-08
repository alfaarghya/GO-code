package main

import "fmt"

func Array() {

	fmt.Println("----- declare arrays -----")
	var arr1 = [5]int{1,11,111,1111,11111} //declare a array with length
	fmt.Println("arr1 >>",arr1, "length >>", len(arr1))
	fmt.Println("----------")
	
	var arr2 = [...]int{1,11,111,1111} //declare a array with out length
	fmt.Println("arr2 >>",arr2, "length >>", len(arr2))
	fmt.Println("----------")
	
	arr3 := [5]int{} //not initialized
	arr4 := [5]int{1,2} //partially initialized
	arr5 := [5]int{1,2,3,4,5} //fully initialized
	fmt.Println("arr3 >>",arr3, "length >>", len(arr3))
	fmt.Println("arr4 >>",arr4, "length >>", len(arr4))
	fmt.Println("arr5 >>",arr5, "length >>", len(arr5))
	fmt.Println("----------")
	
	arr6 := [5]int{1:10,3:55} //initialize Only Specific Elements
	fmt.Println("arr6 >>",arr6, "length >>", len(arr6))
	fmt.Println("----------")
	
	fmt.Println("----- access & change the values -----")
	
	fmt.Println("arr6[3] >>",arr6[3]) //access
	fmt.Println("----------")
	
	arr6[4] = 125 //change 
	fmt.Println("arr6 >>",arr6)
	fmt.Println("----------")

}