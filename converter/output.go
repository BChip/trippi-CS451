package main

import "fmt"

func main() {
	numOne := 5
	numTwo := 10
	sum := numOne + numTwo
	fmt.Println("The sum of num1 and num2 is: ", sum)
	if sum > 20 {
		fmt.Println("The sum is greater than 20")
	}
}
