package main

import "fmt"

func main() {
	age := 180
	switch age {
	case 18:
		fmt.Println("You are 18 years old")
	case 19:
		fmt.Println("You are 19 years old")
	case 20:
		fmt.Println("You are 20 years old")
	default:
		fmt.Println("You are not 18, 19 or 20 years old")
	}
}
