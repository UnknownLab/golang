package main

import (
	"fmt"
	"math/big"
)

func main() {
	xs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(xs))
	x, y := multipleReturn()
	fmt.Println(x, y)
	fmt.Println(variadicAdd(1, 2, 3))
	fmt.Println(variadicAdd(xs...))
	x2 := getClosureX2()
	fmt.Println(x2(4.2))

	nextEven := makeEvenGenerator()

	for i := 1; i <= 5; i++ {
		fmt.Println("Even", nextEven())
	}
	fmt.Println("Factorial", factorial(65))
	fmt.Println("start Big Factorial calculating")
	bigFactorial(big.NewInt(100000))
	fmt.Println("end Big Factorial calculating")
	fmt.Println(yes())
	theRecover()
	theEnd()
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func multipleReturn() (int, int) {
	return 5, 6
}

func variadicAdd(args ...float64) float64 {
	total := 0.0
	for _, v := range args {
		total += v
	}
	return total
}

func getClosureX2() func(float64) float64 {
	closure := func(v float64) float64 {
		return v * 2
	}

	return closure
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func bigFactorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, bigFactorial(n.Sub(x, n)))
}

func yes() (text string) {
	defer func() {
		text = "no"
	}()
	return "yes"
}

func theRecover() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

func theEnd() {
	panic("PANIC")
	str := recover()
	fmt.Println(str)
}
