// main.go
package main

import "fmt"

func main() {
	fmt.Printf("2 + 2 = %v \n", sum(2,2))
	fmt.Printf("2 - 2 = %v \n", sub(2,2))
	fmt.Printf("2 * 2 = %v \n", mult(2,2))
	fmt.Printf("2 / 2 = %v \n", div(2,2))
}

func sum(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}
func mult(a int, b int) int {
	return a * b
}
func div(a int, b int) int {
	return a / b
}
