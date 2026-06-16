package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Usage: linecounter <filename>\nExample: linecounter text.txt")
	}
	file, err := os.OpenFile(args[0], os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("unable to open file ", args[0])
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan() {
		lines++
	}

	fmt.Println(lines)
}
