package main

import "fmt"

type Person struct {
	Talks
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func main() {
	a := new(Android)
	a.Name = "Tolya"
	a.Person.Talk()
}

type Talks interface {
	Talk()
}
