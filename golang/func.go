package main

import "fmt"

func main() {
	result, mesage := sum(1, 2)
	fmt.Printf("Tong cua 2 so = %d \n", result)
	fmt.Printf("%s \n\n\n", mesage)

	count, dogs := getDogs()
	// _, dogs := getDogs() // just get dog array, skip count
	fmt.Printf("Co %d chu cho \n", count)
	fmt.Println(dogs)
}

func sum(a int, b int) (int, string) {
	// var sum int
	sum := a + b
	if a > b {
		return sum, "So thu nhat > So thu 2"
	}
	return sum, "So thu hai > So thu nhat"
}

//Go
func getDogs() (count int, dogs []string) {
	dogs = []string{
		"Pug",
		"Chihuahua",
		"Husky",
		"Pit Bull",
	}
	count = 4
	// return

	// Or return
	return count, dogs
}
