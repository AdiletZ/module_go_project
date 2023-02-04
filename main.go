package main

import "fmt"

type Person struct {
	name    string
	surname string
	age     int
	weight  float64
}

func (p *Person) changeName() {
	p.name = "new_name"
}

func main() {
	some_person := Person{"Adilet", "Zhakenger", 19, 59}
	some_person.changeName()
	fmt.Println(some_person.name)
}
