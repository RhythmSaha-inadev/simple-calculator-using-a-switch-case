package main

import "fmt"

func main() {

	fmt.Println("Enter two no.")
	var a, b int
	fmt.Scan(&a, &b)
	var op string
	fmt.Println("Enter an operator (+, -, *, /)")
	fmt.Scan(&op)

	switch op {
	case "+":
		fmt.Println("Add is ", a+b)
	case "-":
		fmt.Println("Sub is ", a-b)
	case "*":
		fmt.Println("Mul is ", a*b)
	case "/":
		fmt.Println("Div is ", a/b)
	default:
		fmt.Println("Invalid operator")

	}

}
