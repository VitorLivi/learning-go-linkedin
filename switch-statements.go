package main

import (
	"fmt"
	"time"
)

func main() {
	weekday := time.Now().Weekday()
	fmt.Printf("Today is %v\n", weekday)

	dayNumber := int(weekday)
	fmt.Printf("The day as a number is %v\n", dayNumber)

	var result string

	switch dayNumber {
	case 1:
		result = "Monday"
	case 2:
		result = "Tuesday"
	case 3:
		result = "Wednesday"
	case 4:
		result = "Thursday"
	case 5:
		result = "Friday"
	default:
		result = "Weekend"
	}

	fmt.Println(result)
	x := -42
	switch {
	case x < 0:
		result = "Negative"
	case x == 0:
		result = "Equal to zero"
	default:
		result = "Greater than zero"
	}
	fmt.Println(result)
}
