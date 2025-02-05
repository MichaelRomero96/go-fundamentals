package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Print("Type the first number: ")
	fmt.Scan(&a)

	fmt.Print("Type the second number: ")
	fmt.Scan(&b)
	var result = sum(a, b)
	println("The sum result is: ", result)
}

func sum(a int, b int) (result int) {
	result = a + b
	return
}
