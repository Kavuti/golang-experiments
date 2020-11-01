package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	var christian person

	christian.firstName = "Christian"
	christian.lastName = "Cavuti"
	fmt.Println(christian)
	fmt.Printf("%+v\n", christian)
}
