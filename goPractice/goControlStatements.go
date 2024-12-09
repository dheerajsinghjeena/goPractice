package goPractice

import "fmt"

func DecisionMaking() {
	var myNumber int = 23

	if myNumber != 53 {
		fmt.Printf("%d is not equal to 53", myNumber)
	} else {
		fmt.Printf("%d is equal to 34", myNumber)
	}
	// like there are more of it such as if else if, nested if and more
}

func ForLoop() {
	// var loopingNum int = 5

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(loopingNum * i)
	// }

	// myLoopingArray := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// for _, j := range myLoopingArray {
	// 	fmt.Println(j)
	// }

	// for loop for string
	// myString := "This is learning phase of GoLang"
	// for _, j := range myString {
	// 	fmt.Println(string(j))
	// }

	// using map method
	// myMap := map[string]string{
	// 	"Name":       "Dheeraj Singh Jeena",
	// 	"Age":        "Twenty Two",
	// 	"Profession": "Node JS Developer",
	// }

	// for i, j := range myMap {
	// 	fmt.Println(i, j)
	// }

	// // map with different types of values
	// myNewMap := map[string]interface{}{
	// 	"Name":       "Dheeeraj Singh Jeena",
	// 	"Age":        34,
	// 	"Address":    "Delhi",
	// 	"Experience": 0.6,
	// }

	// for i, j := range myNewMap {
	// 	fmt.Println(i, j)
	// }
	// // simply accessing map in go
	// simpleMapObject, ok := myNewMap["Name"].(string)
	// if ok {
	// 	fmt.Println(simpleMapObject)
	// }

	// print a table of number
	tableNumber := 23

	for i := 1; i <= 10; i++ {
		tableNumber += i
	}

	fmt.Println("got my number: ", tableNumber)
}
