package main

import "fmt"

func Slice()  {
  fmt.Println("----- declare slice -----")
	var sli1 = []int{1,11,121,1221,11211} //declare a slice
	fmt.Println("sli1 >>",sli1, "length >>", len(sli1), "capacity >>", cap(sli1))
	fmt.Println("----------")
	
	var arr1 = [6]int{10,20,30,40,50,60}
  var sli2 = arr1[3:6] //create a slice with an array
	fmt.Println("sli2 >>",sli2, "length >>", len(sli2),"capacity >>", cap(sli2))
	fmt.Println("----------")
 
	var sli3 = make([]int, 5,10) //create a slice make() function
	fmt.Println("sli3 >>",sli3, "length >>", len(sli3), "capacity >>", cap(sli3))
	fmt.Println("----------")

  fmt.Println("----- Append into slice -----")
  var sli4 = []int{2,4,6,8}
  fmt.Println("sli4 >>",sli4, "length >>", len(sli4), "capacity >>", cap(sli4))
  sli4 = append(sli4, 10, 20) // append new elements
  fmt.Println("sli4 >>",sli4, "length >>", len(sli4), "capacity >>", cap(sli4))
  sli4 = append(sli4, sli1...) //append another slice
  fmt.Println("sli4 >>",sli4, "length >>", len(sli4), "capacity >>", cap(sli4))
  fmt.Println("----------")



}
