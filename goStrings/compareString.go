package goStrings

import (
	"fmt"
	"strings"
)

// in this we will learn about the different method to compare strings
func CompareString() {
	myString1 := "Hello World"
	myString2 := "Hello World"
	myString3 := "hello world"
	myString4 := "hello world"

	// first comparing string using simple comparison operator
	// use case of different comparison operator
	// == will check if both string match a == b it's false
	// != will check the if both not match a != b it's true
	// < will check the left hand side string should be less than the right hand side in lexicographical order(dictionary order)  a < d it's true
	// > will check the left hand side string should be greater than the right hand side in lexicographical order(dictionary order) d > a it's also true

	var checkStringMatch bool = myString1 == myString2
	var checkStringMatchSecond bool = myString3 == myString4

	fmt.Printf("My first string match is: %t\n", checkStringMatch)
	fmt.Printf("My second string match is: %t\n", checkStringMatchSecond)

	// now we will use strings.compare function, this function also compare two strings based on lexicographical order and return
	// 0 if string1 is equal to string2
	// 1 if string1 is lexicographically greater than string2
	// -1 if string1 is lexicographical less than string2

	var stringCompareMethod int = strings.Compare(myString1, myString2)
	fmt.Println(stringCompareMethod, " got my method to compare string") // return 0 as myString1 is lexicographical equal to myString2
}
