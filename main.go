package main

import "fmt"

func main() {
	var operator string
	fmt.Println("Enter operator: +, -, *, /")
	_, err := fmt.Scan(&operator)
	if err != nil {
		fmt.Println(err)
	}
	var num1, num2 float64
	fmt.Println("Enter two numbers: ")
	_, err = fmt.Scan(&num1, &num2)
	if err != nil {
		fmt.Println(err)
	}
	switch operator {
	case "+":
		fmt.Println(num1 + num2)
	case "-":
		fmt.Println(num1 - num2)
	case "*":
		fmt.Println(num1 * num2)
	case "/":
		fmt.Println(num1 / num2)
	default:
		fmt.Println("Invalid operator")
	}

}
func addition(a float64, b float64) float64 {
	return a + b
}
