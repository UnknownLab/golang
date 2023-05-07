package main

import "fmt"

func main() {
	for_f()
	for_f2()
	for_if_f()
	switch_f()
}

func for_f() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
}

func for_f2() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func for_if_f() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}

func switch_f() {
	var i = 4
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}
}
