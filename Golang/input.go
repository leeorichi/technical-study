package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var limitNumber int = 3
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Nhập %d số tự nhiên, cách nhau bởi dấu cách: ", limitNumber)
	inputStr, _ := reader.ReadString('\n')

	inputArrStr := strings.Split(strings.TrimSpace(inputStr), " ")

	if len(inputArrStr) != limitNumber {
		fmt.Printf("Bạn phải nhập đúng %d số tự nhiên!", limitNumber)
		return
	}

	sum := 0
	for _, str := range inputArrStr {
		num, err := strconv.Atoi(str)
		if err == nil {
			sum += num
		}
	}

	fmt.Printf("Tổng của %d số tự nhiên là: %d", limitNumber, sum)
}
