package goPractice

import "fmt"

func Anonymous() {
	var num1, num2 int = 12, 33
	myValue := func(num1 int, num2 int) int {
		myResult := num1 + num2 // Declare myResult here
		return myResult
	}(num1, num2)

	fmt.Println("hence result for myValue function is: ", myValue)
}
