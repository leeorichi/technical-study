package main

import "fmt"

func main() {

	// case1_array_slice()

	// case2_slice_method()

	// case3_append_method()

	case4_mapping_method()

}

func case1_array_slice() { // default is void function

	//This is Array
	// person := [4]string{
	// 	"LONG 1",
	// 	"LONG 2",
	// 	"LONG 3",
	// 	"LONG 4",
	// }

	// This is Slice, # vs array is not number of size
	person := []string{
		"LONG 1",
		"LONG 2",
		"LONG 3",
		"LONG 4",
		"LONG 5",
		"LONG 6",
		"LONG ...",
	}

	fmt.Println("My Name is " + person[1] + " with my best friend is " + person[2]) // My Name is LONG 2 with my best friend is LONG 3
}

func case2_slice_method() {

	person := []string{
		"LONG 1",
		"LONG 2",
		"LONG 3",
		"LONG 4",
		"LONG 5",
		"LONG 6",
		"LONG ...",
	}

	new_person := person[2:]
	fmt.Println(new_person) // print [LONG 3 LONG 4 LONG 5 LONG 6 LONG ...]
}

func case3_append_method() {
	message := []string{
		"My", "Name", "is",
	}
	message = append(message, "Long", "Ha", "And", "I'am", "a", "Senior Developer")
	fmt.Println(message)
}

func case4_mapping_method() {
	/* IN PHP , I wanna to array with:
	   $myInfo = [
	       'name' => 'LongHa',
	       'age' => "28",
	       'job' => 'Software Engineer',
	   ];
	*/
	myInfo := map[string]string{
		"name": "LongHa",
		"age":  "28",
		"job":  "Software Engineer",
	}

	fmt.Printf(myInfo["name"])
}
