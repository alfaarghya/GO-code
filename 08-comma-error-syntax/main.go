package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(" Enter your input >")
	reader := bufio.NewReader((os.Stdin))
	input, _ := reader.ReadString(('\n'))
	fmt.Println(" >>", input)

}
