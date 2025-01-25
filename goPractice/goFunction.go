package goPractice

import "fmt"

func multiply() int {
	var num1, num2 int = 12, 33

	resultant := num1 * num2

	return resultant
}

// parameterized function
func parameterizedFunc(num1, num2 int) {
	firstNum := num1
	secondNum := num2

	getResult := firstNum + secondNum

	fmt.Println("Got my result as: ", getResult)
}

func callByValue(num1 int) int {
	num1 = num1 + 10
	return num1
}

func callByReference(num1 *int) int {
	*num1 = *num1 + 10
	return *num1
}

func callVariadicFunc(numbers ...int) int {
	// fmt.Println("My multiple parameter value is: ", numbers)
	total := 0
	for _, j := range numbers {
		total += j
	}
	return total
}

func FunctionUse() {
	// fmt.Println("hello world");
	var myNumber = 12
	var myFuncResult = multiply()

	fmt.Println("My result for multiply is: ", myFuncResult)

	// calling parametrized function
	parameterizedFunc(12, 44)

	// call by value function (original value will not be affected on passing to another function)
	fmt.Println("Before passing value: ", myNumber)
	var newVal = callByValue(myNumber)
	fmt.Println("modified value is: ", newVal)
	fmt.Println("After passing value: ", myNumber)

	// call by reference (original value will be affected on passing value to another function)
	fmt.Println("Value before call by reference: ", myNumber)
	callByReference(&myNumber)
	fmt.Println("Value after call by reference: ", myNumber)

	// use of Variadic function (this allow to pass unknown number of parameters to my function)
	fmt.Println("first passing three values: ")
	callVariadicFunc(12, 3, 4, 23, 3)
}

func PrintSum() int {
	var num1, num2 int = 12, 12
	return AddNum(num1, num2)
}
