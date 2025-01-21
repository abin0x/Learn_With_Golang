package main

import "fmt"

func main() {

	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	var num1, num2 int
	fmt.Println("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&num2)
	sum := num1 + num2
	println("Hello ", name)
	fmt.Println("Sum is: ", sum)
	fmt.Println("Thank you for using our program")
	fmt.Println("Have a nice day")

}
