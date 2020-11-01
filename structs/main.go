package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Morrison",
		contact: contactInfo{
			email: "jim.morrison@email.com",
			zip:   00000,
		},
	}

	jim.print()
	jim.updateName("John")
	jim.print()
}

func (p *person) updateName(name string) {
	p.firstName = name
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
