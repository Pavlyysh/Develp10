package main

import "fmt"

func main() {
	x := []int{1, 2, 3}

	sl := Map(x, func(num int) int { return num + 2 })
	fmt.Println(sl)

	sl = Filter(x, func(num int) bool { return num%2 == 0 })
	fmt.Println(sl)

	res := Reduce(x, func(a, b int) int { return a + b }, 0)
	fmt.Println(res)
}

func Map[T, U any](
	slice []T,
	calc func(T) U,
) []U {
	result := make([]U, len(slice))
	for i, item := range slice {
		result[i] = calc(item)
	}

	return result
}

func Filter[T any](
	slice []T,
	filter func(T) bool,
) []T {
	result := make([]T, 0, len(slice))

	for _, item := range slice {
		if filter(item) {
			result = append(result, item)
		}
	}

	return result
}

func Reduce[T, U any](slice []T, accumulator func(T, U) U, initial U) U {
	result := initial

	for _, item := range slice {
		result = accumulator(item, result)
	}

	return result
}
