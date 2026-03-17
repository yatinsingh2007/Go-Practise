package main

import (
	"fmt"
	"time"
)
func main () {
	fmt.Println(time.Now())

	formattedTime := time.Now()
	fmt.Println(formattedTime.Format("2006-01-02 15:04:05"))
	
}
