package main

import "fmt"

// type Animal struct{}

type Animal interface {
	Speak() string
	MoveBy() string
	Eat() string
	Tail() bool
	vvv() string
}

// method
type Dog struct{ Name string }
type Chicken struct{ Name string }

func (d Dog) Speak() string {
	return "Gogogoogogogo"
}

func (d Dog) MoveBy() string { return "4 legs" }
func (d Dog) Eat() string    { return "Rice, sh!t" }
func (d Dog) Tail() bool     { return true }

func (d Chicken) Speak() string  { return "o` o` o` ó ó ooooo" }
func (d Chicken) MoveBy() string { return "2 Legs" }
func (d Chicken) Eat() string    { return "Rice" }
func (d Chicken) Tail() bool     { return true }

func main() {

	dog1 := Dog{
		Name: "cho con 1",
	}

	fmt.Println(dog1.Speak())

	ckn1 := Chicken{
		Name: "Ga U muoi",
	}

	fmt.Println(ckn1.Speak())
}
