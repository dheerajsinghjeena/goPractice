package goPractice

import "fmt"

// second function
func secondFunc() {
	fmt.Println("Second initial call")
}

func thirdFunc() {
	fmt.Println("Third initial call")
}

func DeferKeyword() {
	// initial function
	// a defer keyword in go is a LIFO i.e. last in first out
	// so if we define a defer keyword then that function will be executed at last
	// so if I have multiple defer keyword then the last defer keyword will be executed first
	func() {
		fmt.Println("First initial call")
	}()

	//second function
	defer secondFunc()

	// defer function
	defer thirdFunc()
}
