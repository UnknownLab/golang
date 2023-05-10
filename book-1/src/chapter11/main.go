package main

import (
	"fmt"
	mp "mypackage/math"
)

// go mod init mypackage
func main() {
	xs := []float64{1, 2, 3, 4}
	avg := mp.Average(xs)
	fmt.Println(avg)
}
