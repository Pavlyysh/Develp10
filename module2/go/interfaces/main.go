package main

import "fmt"

func main() {
	var x int64
	num := 10
	Describe(num)
	st := "hello"
	Describe(st)

	Describe(x)
}

func Describe(a any) {
	fmt.Printf("Type: %T\nValue: %v\n", a, a)
}
