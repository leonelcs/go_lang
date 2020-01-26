package main

import "fmt"

func main() {
	array1 := []int{1, 2, 3, 4}
	array2 := []int{5, 6, 7, 8}

	arrayMerged := append(array1, array2...)

	fmt.Println(arrayMerged)
}
