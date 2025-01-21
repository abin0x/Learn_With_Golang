package main

import "fmt"

func sum(num1 int, num2 int) {
	add := num1 + num2
	fmt.Println("Sum is: ", add)
}

func main() {
	a := 10
	b := 20
	sum(a, b)
}
