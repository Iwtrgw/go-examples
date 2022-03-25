package main

import "fmt"

func main() {
	array := [5]int64{}
	fmt.Println(array)

	array[0] = 8
	array[1] = 9
	array[2] = 7
	fmt.Println(array)
	fmt.Println(array[2])
}
