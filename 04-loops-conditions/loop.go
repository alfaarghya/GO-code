package main

import ("fmt")

func Loop() {
  for i := 0; i < 10; i++ {
    fmt.Println(i)
  }
  
  fmt.Println("----------")
}

func NestedLoop() {
  for i := 1; i <= 5; i++ {
    for j := 1; j <= i; j++ {
      fmt.Printf("*")
    }
    fmt.Println()
  }

  fmt.Println("----------")
}

func Continue() {
  for i := 10; i > 0; i-- {
    if i == 5 {
      continue
    }
    fmt.Println(i)
  }

  fmt.Println("-----------")
}

func Break() {
  for i := 10; i > 0; i-- {
    if i == 5 {
      break;
    }

    fmt.Println(i)
  }

  fmt.Println("----------")
}

func Range() {
  colors := [5]string{"white", "black", "red", "green", "blue"}; //array of string

  for idx, val := range colors {
    fmt.Println(idx, "->", val)
  }

  fmt.Println("----------")
}
