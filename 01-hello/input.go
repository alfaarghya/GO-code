package main

import "fmt"

func Input() {
	var name string //variable

	fmt.Println("What is your name?")
	fmt.Print("> ")
	fmt.Scan(&name) //taking input from user

	fmt.Println(">> Greetings, "+name)

}