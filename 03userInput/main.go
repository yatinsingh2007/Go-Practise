package main


import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Enter your rating of pizza :")
	reader := bufio.NewReader(os.Stdin)
	// comma ok , comma error
	input , _ := reader.ReadString('\n')
	fmt.Printf("The rating of your pizza is : %s\n", input)
	fmt.Printf("The type of rating of your pizza is : %T\n", input)
}