package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "it's a content, that I want to write to a file"
	writeFile("./newFile.txt", content)
	readFile("./newFile.txt")

}

func writeFile(fileName string, content string) {
	file, err := os.Create(fileName) //create a file

	if err != nil { //if catch the error
		panic(err) //stop the execution
	}

	length, err := io.WriteString(file, content) //write the content to the file

	if err != nil { //if catch the error
		panic(err) //stop the execution
	}

	fmt.Println("length >> ", length)

	defer file.Close() //close the file
}

func readFile(fileName string) {
	databyte, err := os.ReadFile(fileName)
	if err != nil { //if catch the error
		panic(err) //stop the execution
	}

	fmt.Println(string(databyte))
}
