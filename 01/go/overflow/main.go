package main

import "fmt"

func main() {
	overflow(12)
}

func overflow(x int8) {

	for {
		fmt.Println(x)
		x++
	}
}
