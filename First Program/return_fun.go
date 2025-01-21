package main

import "fmt"

func sum(a int, b int) int {
	c := a + b
	return c
}
func main() {
	a := 11
	b := 22
	c := sum(a, b)
	fmt.Println("Sum is: ", c)
}
