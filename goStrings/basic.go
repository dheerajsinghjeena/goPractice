package goStrings

import "fmt"

// in this we will learn basic string use case in go lang
// there are two type of string or string literals first is using double quotes and second is using single quotes single quote string is not supported in go lang.
// double quotes("") -> it support using of escape character but when using double quotes then we need to define our string in single line
// backtick quotes(``) -> it does not support using of escape character but support multi line string declaration.

// in go string are immutable
func BasicOfStrings() {
	var myString string = `dheeraj is good boy`

	for index, char := range myString {
		result := fmt.Sprintf(`%c is my string and it's index is %d\n`, char, index)
		fmt.Println(result)
	}

	// find the length of a string
	fmt.Printf("length of string is %d", len(myString))
}
