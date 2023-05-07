package main

import "fmt"

func main() {
	array_f()
	array_f2()
	cast_f()
	slice_f()
	slice_f2()
	slice_append_f()
	slice_copy_f()
	map_f()
	map_2_f()
}

func array_f() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
}

func array_f2() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += x[i]
	}
	fmt.Println(total / 5)
}

func cast_f() {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	var total float64 = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x)))
}

func slice_f() {
	x := make([]float64, 5, 10)
	fmt.Println(x)
}

func slice_f2() {
	arr := [5]float64{1, 2, 3, 4, 5}
	x := arr[0:5]
	fmt.Println(x)
}

func slice_append_f() {
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

func slice_copy_f() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}

func map_f() {
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
}

func map_2_f() {
	{
		elements := map[string]map[string]string{
			"H": map[string]string{
				"name":  "Hydrogen",
				"state": "gas",
			},
			"He": map[string]string{
				"name":  "Helium",
				"state": "gas",
			},
			"Li": map[string]string{
				"name":  "Lithium",
				"state": "solid",
			},
			"Be": map[string]string{
				"name":  "Beryllium",
				"state": "solid",
			},
			"B": map[string]string{
				"name":  "Boron",
				"state": "solid",
			},
			"C": map[string]string{
				"name":  "Carbon",
				"state": "solid",
			},
			"N": map[string]string{
				"name":  "Nitrogen",
				"state": "gas",
			},
			"O": map[string]string{
				"name":  "Oxygen",
				"state": "gas",
			},
			"F": map[string]string{
				"name":  "Fluorine",
				"state": "gas",
			},
			"Ne": map[string]string{
				"name":  "Neon",
				"state": "gas",
			},
		}

		delete(elements, "Ne")

		if el, ok := elements["Li"]; ok {
			fmt.Println(el["name"], el["state"])
		}
	}
}
