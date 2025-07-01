package main

import (
	"fmt"
	"math"
)

func main() {
	i1, i2, i3 := 1, 2, 3
	intSum := i1 + i2 + i3
	fmt.Println(intSum)

	f1, f2, f3 := 1.1, 2.2, 3.3
	floatSum := f1 + f2 + f3
	fmt.Println(floatSum)

	total := float64(intSum) + floatSum
	fmt.Println("Total: ", total)

	product := float64(intSum) * floatSum
	fmt.Println("Product: ", product)

	preciseNumber := math.Round(product*100) / 100
	fmt.Println("Precise Number: ", preciseNumber)

	fmt.Println("The value of PI: ", math.Pi)

	circleRadius := 15.5
	circumReference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference: %.3f\n", circumReference)
}
