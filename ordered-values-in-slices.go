package main

import (
	"fmt"
	"sort"
)

func main() {
	// var colors = []string{"Red", "Green", "Blue"}
	// fmt.Println(colors)

	var colors = make([]string, 0, 3)
	colors = append(colors, "Red", "Green", "Blue")
	fmt.Println(colors)
	colors = append(colors, "Purple", "Yellow", "Orange")
	fmt.Println(colors)

	colors = remove(colors, 0)
	fmt.Println(colors)

	sort.Strings(colors)
	fmt.Println(colors)
}

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}
