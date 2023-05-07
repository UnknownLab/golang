package main

import "fmt"

var b = "Hello World"

func main() {
	var x string = "Hello World"
	var y string
	y = "Hello World 2"
	fmt.Println(x)
	fmt.Println(y)
	x = "Hello World 3"
	fmt.Println(x)
	fmt.Println(x == y)
	y = "Hello World 3"
	fmt.Println(x == y)

	z := "Hello World 4"
	fmt.Println(z)
	var a = "Hello World"
	fmt.Println(a)
	fmt.Println(b)
	f()
	f2()
}

func f() {
	fmt.Println(b)
	const x = "Hello World"
	fmt.Println(x)
	// error
	//x = "Some other string"

	var (
		a = 5
		b = 10
		c = 15
	)

	fmt.Println(a, b, c)
}

func f2() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
