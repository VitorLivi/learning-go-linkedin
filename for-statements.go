package main

import (
	"fmt"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}

	for i := 0; i < len(colors); i++ {
		color := colors[i]
		fmt.Println(color)
	}

	fmt.Print("---------------------------------------\n\n")

	for i := range colors {
		color := colors[i]
		fmt.Println(color)
	}

	fmt.Print("---------------------------------------\n\n")

	for _, color := range colors {
		fmt.Println(color)
	}

	fmt.Print("---------------------------------------\n\n")

	states := make(map[string]string)
	states["WA"] = "Washington"
	states["CA"] = "California"
	states["NY"] = "New York"

	for state, _ := range states {
		fmt.Println(states[state])
	}

	fmt.Print("---------------------------------------\n\n")

	value := 0
	sum := 0
	for value < 5 {
		sum += value
		fmt.Printf("Value: %v\n", value)
		fmt.Printf("Sum: %v\n", sum)
		value++
	}

	sum = 1
	for sum < 1000 {
		sum += sum
		if sum > 200 {
			goto theEnd
		}
	}

theEnd:
	println("end of program")
	fmt.Printf("Sum: %v\n", sum)
}
