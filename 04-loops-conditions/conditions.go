package main

import (
	"fmt"
)

func IfElse() {
  var age int;

  fmt.Printf("What's your age? \n>")
  fmt.Scan(&age)
  
  if(age == 18) {
    fmt.Println(">> You are ready to get your Driver License")
  } else if(age > 18) {
    fmt.Println(">> You are eligible for a Driver License")
    } else {
    fmt.Println(">> Too young to Drive!! Come back when you are 18 years old")
  }
  
  fmt.Println("----------")
  
}

func SwitchSingleCase() {
  var day int;
  
  fmt.Printf("Give Day Number(1 to 7)? \n>")
  fmt.Scan(&day)
  
  switch day {
  case 1:
    fmt.Println(">> Sunday")
  case 2:
    fmt.Println(">> Monday")
  case 3:
    fmt.Println(">> Tuesday")
  case 4:
    fmt.Println(">> Wednesday")
  case 5:
    fmt.Println(">> Thursday")
  case 6:
    fmt.Println(">> Friday")
  case 7:
    fmt.Println(">> Saturday")
  default:
    fmt.Println(">> Not a day!!!!")
  }
  
  fmt.Println("----------")
}

func SwitchMultiCase() {
  var day int;
  
  fmt.Printf("Give Day Number(1 to 7)? \n>")
  fmt.Scan(&day)
  
  switch day {
  case 2,3,4,5,6:
    fmt.Println(">> Week Day -> means work day")
  case 1,7:
    fmt.Println(">> Weekend -> means enjoy day")
  default:
    fmt.Println(">> Not a day!!!!")
  }
  
  fmt.Println("----------")
}
