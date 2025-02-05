package main

import "fmt"

func main() {
	var revenue float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	fmt.Println(revenue, taxRate)
}
