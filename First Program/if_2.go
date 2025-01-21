package main

import "fmt"

func main() {
	age := 13
	sex := "F"
	if age < 18 && sex == "M" {
		fmt.Println("You are eligible to vote")
	} else {
		fmt.Println("You are not eligible to vote")
	}
}
