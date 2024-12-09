package goPractice

import (
	"fmt"
)

func MyFunction() {
	fmt.Println("hello world")
}

func ArithmeticOperators() {
	// addition operator start
	var resultantSum int
	num1, num2 := 12, 34
	resultantSum = num1 + num2

	fmt.Println("Sum of two number is: ", resultantSum)

	// subtraction operator start
	var resultantSub int
	subNum1, subNum2 := 10, 20

	resultantSub = subNum1 + subNum2
	fmt.Println("Sub of two number is: ", resultantSub)

	// multiplication operator start
	var resultantMulti int
	multiNum1, multiNum2 := 11, 23

	resultantMulti = multiNum1 * multiNum2

	fmt.Println("Multiplication of Two number is: ", resultantMulti)

	// division of operator
	var resultantDivision float64
	diviNum1, diviNum2 := 23.0, 12.0

	resultantDivision = diviNum1 / diviNum2

	fmt.Println("Division of number 1 and number 2 is: ", resultantDivision)

	// modulus operator don't work with floating number
	var resultantModulus int
	modNum1, modNum2 := 13, 33

	resultantModulus = int(modNum2) % int(modNum1)
	fmt.Println("Modulus of two number is: ", resultantModulus)
}

func RelationalOperator() {
	// using equal to operator ==
	num1, num2 := 12, 22

	var result1 = num1 == num2
	fmt.Println("Result For Equal To Operator", result1)

	// using not equal to operator !=
	var result2 = num1 != num2
	fmt.Println("Result for Not Equal: ", result2)

	// using greater to operator >
	var result3 = num1 > num2
	fmt.Println("Result for Greater than: ", result3)

	// using less than operator <
	var result4 = num1 < num2
	fmt.Println("Result for Less than: ", result4)

	// using greater than equal to
	var result5 = num1 >= num2
	fmt.Println("Result for greater than equal to: ", result5)

	// using less than equal to
	var result6 = num1 <= num2
	fmt.Println("Result for less than equal to: ", result6)
}

func LogicalOperators() {
	// logical && operator
	var num1, num2 int = 11, 34

	if num1 != 0 && num2 != 12 {
		fmt.Println("True")
	}

	if num1 == 3 {
		fmt.Println("True")
	}

	if num1 == 11 || num2 == 23 {
		fmt.Println("True")
	}

	if num1 != 11 {
		fmt.Println("True")
	}
}
