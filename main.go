package main

import (
	"fmt"
	"newSetupGo/goStrings"
	"time"
)

func main() {
	// Basic of strings
	// goStrings.BasicOfStrings()
	// goStrings.CompareString()
	fmt.Printf("\nStarted the execution @ %v \n", time.Now())
	var rawdata = goStrings.GoAssignment{}.LoadRawDataList()
	goStrings.GoAssignment{}.DoTheWork(rawdata)
	fmt.Printf(" \n ")
	fmt.Printf("Finished the execution @ %v ", time.Now())
}
