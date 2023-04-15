package main

import "fmt"

func main() {
	var s string
	fmt.Println("nhap day 3 chu so")

	fmt.Scanf("%s", &s)
	Var_dump(s)
}

func Var_dump(expression ...interface{}) {
	fmt.Println(fmt.Sprintf("%#v", expression))
}
