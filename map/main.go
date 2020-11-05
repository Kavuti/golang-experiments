package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	for key, value := range colors {
		fmt.Printf("%s: %s\n", key, value)
	}
}
