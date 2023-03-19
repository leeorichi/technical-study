package main

import (
	"fmt"
	// "strconv"
	// "strings"
)

func main() {
	var input string
	fmt.Print("Nhập 10 số tự nhiên: ")
	fmt.Scanln(&input)

	// Chuyển chuỗi input thành slice các số nguyên
	// numbersStr := strings.Fields(input)
	fmt.Println(input)
	// numbers := make([]int, len(numbersStr))
	// for i, s := range numbersStr {
	// 	num, _ := strconv.Atoi(s)
	// 	numbers[i] = num
	// }

	// // Bubble sort
	// for i := 0; i < len(numbers)-1; i++ {
	// 	for j := i + 1; j < len(numbers); j++ {
	// 		if numbers[i] > numbers[j] {
	// 			numbers[i], numbers[j] = numbers[j], numbers[i]
	// 		}
	// 	}
	// }

	// fmt.Println(numbers)
}