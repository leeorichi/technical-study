package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("[+] Hello Loopy!")

	// basic_loop()

	// while_loop()

	foreach_loop()

}

func basic_loop() {
	sum := 0
	count := 10
	for i := 0; i < count; i++ {
		// fmt.Println(i)
		sum += 1
	}
	s := strconv.Itoa(sum) // convert int to string
	fmt.Println("(For Loop) SUM = " + s)
	// fmt.Printf("SUM = %d ", sum) // with direct output %d
}

func while_loop() {
	sum := 0
	for sum < 1000 {
		sum += 1
		fmt.Println(sum)
	}
	fmt.Printf("(Loop while case) SUM = %d \n", sum) // with direct output %d
}

func foreach_loop() {
	//init slice of array elements;
	bonus := map[string]int{
		"day 1": 100,
		"day 2": 110,
		"day 3": 140,
		"day 4": 140,
		"day 5": 10,
		"day 6": 99,
	}

	for key, value := range bonus {
		fmt.Printf("In a %s,I Have earned $%d \n", key, value) // with direct output %d
		// fmt.Println("In a " + key + "I earned" + value)
	}
}
