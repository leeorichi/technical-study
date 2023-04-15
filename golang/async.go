package main

import "time"

func main() {

	println("Sync Mode:")
	function(1)
	function(2)
	function(3)

	println("Async Mode:")
	go function(1)
	go function(2)
	go function(3)

	time.Sleep(time.Second)
}

func function(number int) {
	println(number)
}
