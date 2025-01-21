package main

import "fmt"

// Function to get user name
func getUserName() string {
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanln(&name)
	return name
}

// Function to get two numbers from the user
func getTwoNumbers() (int, int) {
	var num1, num2 int
	fmt.Println("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter second number: ")
	fmt.Scanln(&num2)
	return num1, num2
}

// Function to calculate the sum of two numbers
func calculateSum(a int, b int) int {
	return a + b
}

// Function to display a greeting message
func displayGreeting(name string) {
	fmt.Println("Hello", name)
}

// Function to display the result
func displayResult(sum int) {
	fmt.Println("Sum is:", sum)
}

// Function to display a thank-you message
func displayThankYouMessage() {
	fmt.Println("Thank you for using our program")
	fmt.Println("Have a nice day")
}

func main() {
	// Get user input
	name := getUserName()
	num1, num2 := getTwoNumbers()

	// Perform calculation
	sum := calculateSum(num1, num2)

	// Display messages
	displayGreeting(name)
	displayResult(sum)
	displayThankYouMessage()
}
