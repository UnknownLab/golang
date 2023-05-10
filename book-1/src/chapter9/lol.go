package main

type Person2 struct {
	Talks2
	Name string
}

//func (p *Person) Talk() {
//	fmt.Println("Hi, my name is", p.Name)
//}

type Android2 struct {
	Person2
	Model string
}

func main() {
	a := new(Android2)
	a.Name = "Tolya"
	// LOL moment, compile without errors !!!
	a.Talk()
}

type Talks2 interface {
	Talk()
}
