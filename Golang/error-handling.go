package main

import "errors"

func main() {
	// if m, e := helloPerson("LongHa"); e != nil {

	if m, e := helloPerson(""); e != nil {
		println("Error: " + e.Error())
	} else {
		println(m)
	}
}

func helloPerson(name string) (string, error) {
	if name == "" {
		return "", errors.New("No name given")
	}
	return "Hello, My name is " + name, nil
}
