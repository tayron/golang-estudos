package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	checkNumber(getValue())
}

func getValue() string {

	if len(os.Args)-1 == 0 {
		fmt.Println("One arg is required for this program")
		os.Exit(1)
	}

	value := os.Args[1]
	return value
}

func checkNumber(value string) {
	fValue, err := strconv.ParseFloat(value, 64)

	if err != nil {
		fmt.Printf("The value <%s> is not a number \n", value)
		os.Exit(1)
	}

	fmt.Printf("The value <%f> is a valid number \n", fValue)
}
