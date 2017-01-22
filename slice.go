package main

// import "fmt"

// func main() {
// 	// Better alternative of Arrays
// 	// var numbers []int
// 	// printSlice(numbers)
// 	// numbers = append(numbers, 10)
// 	// printSlice(numbers)
// 	// numbers = append(numbers, 30)
// 	// numbers = append(numbers, 40)
// 	// numbers = append(numbers, 50)
// 	// numbers = append(numbers, 60)
// 	// printSlice(numbers)
// 	// //Removing data
// 	// numbers = append(numbers[2:3])

// 	// printSlice(numbers)
// 	// How to increase the size of the slice.
// 	slice := make([]int, 10, 15)
// 	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
// 	newSlice := make([]int, len(slice), 2*cap(slice))
// 	for i := range slice {
// 		newSlice[i] = slice[i]
// 	}
// 	slice = newSlice
// 	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
// 	//

// }

// func printSlice(x []int) {
// 	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
// }
