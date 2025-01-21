package main

import "fmt"

func sum(a int, b int) (int, int) {
	c := a + b
	d := a - b
	return c, d
}
func main() {
	a := 11
	b := 22
	c, d := sum(a, b)
	fmt.Println("Sum is: ", c)
	fmt.Println("Sub is: ", d)
}
