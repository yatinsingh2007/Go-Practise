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
	input , err := reader.ReadString('\n');
	if err != nil{
		fmt.Println(err)
	}
	fmt.Printf("The rating of your pizza is : %s\n", input)
	fmt.Printf("The type of rating of your pizza is : %T\n", input)
}