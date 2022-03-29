package main

import "fmt"

func main() {
	spending := []float64{9.5, 8.0, 10.2, 7.4}
	fmt.Println(spending)
	fmt.Println(len(spending))
	fmt.Println(cap(spending))

	spending = append(spending, 7.0)
	fmt.Println(len(spending))
	fmt.Println(cap(spending))
}
