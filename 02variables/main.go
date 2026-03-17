package main

import "fmt"

func main() {
	// string
	var name string = "Yatin Singh"
	fmt.Println(name)
	fmt.Printf("Variable is of type : %T \n", name)

	// int
	var age int = 21
	fmt.Println(age)
	fmt.Printf("Variable is of type : %T \n", age)

	// bool
	var isStudent bool = true
	fmt.Println(isStudent)
	fmt.Printf("Variable is of type : %T \n", isStudent)

	// float64
	var height float64 = 5.98974937312345645678
	fmt.Println(height)
	fmt.Printf("Variable is of type : %T \n", height)

	// float32
	var weight float32 = 5.9897493732345
	fmt.Println(weight)
	fmt.Printf("Variable is of type : %T \n", weight)

	// Using walrus operator
	bankai := "Bankai"
	fmt.Println(bankai)
	fmt.Printf("Variable is of type : %T \n", bankai)
}