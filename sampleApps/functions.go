package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	readNum1 := bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter the first number:")
// 	str1, _ := readNum1.ReadString('\n')
// 	num1, err := strconv.ParseFloat(strings.TrimSpace(str1), 64)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Value of the number is:", num1)
// 	}
// 	readNum2 := bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter the second number:")
// 	str2, _ := readNum2.ReadString('\n')
// 	num2, err := strconv.ParseFloat(strings.TrimSpace(str2), 64)
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println("Value of the number is:", num2)
// 	}
// 	ret := max(num1, num2)
// 	fmt.Println("The maximum of the two is :", ret)
// }

// func max(num1, num2 float64) float64 {
// 	max := 0.1
// 	if num1 > num2 {
// 		max = num1
// 	} else {
// 		max = num2
// 	}
// 	return max
// }
