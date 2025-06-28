package main

import (
	"fmt"
)

func main() {
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy dog"
	aNumber := 42

	fmt.Println(str1, str2, str3)
	stringLength, err := fmt.Println("The value is", aNumber)

	if err == nil {
		fmt.Println("String length:", stringLength)
	} else {
		fmt.Println("Error:", err)
	}

	fmt.Printf("Value of number: %v\n", aNumber)
	fmt.Printf("Data type: %T\n", aNumber)
}
