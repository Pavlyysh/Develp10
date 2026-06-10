package main

import (
	"fmt"
)

func main() {
	fmt.Println(typeAssert("hello"))
	fmt.Println(typeAssert(12))
	fmt.Println(typeAssert(12.2))

}

func typeAssert(x any) string {
	switch x.(type) {
	case string:
		s := x.(string)
		return fmt.Sprintf("This is a string %s", s)
	case int:
		return fmt.Sprintf("This is an integer %d", x)
	default:
		return "unknown type"
	}
}
