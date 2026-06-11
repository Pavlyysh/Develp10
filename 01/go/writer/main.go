package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

type CustomWriter struct {
	Prefix string
}

type ColorWriter struct {
	Color int
}

func (col *ColorWriter) Write(p []byte) (int, error) {
	switch col.Color {
	case 1:
		color.Red(string(p))
	case 2:
		color.Green(string(p))
	case 3:
		color.Blue(string(p))
	default:
		color.White(string(p))
	}

	return len(p), nil
}

func (custom *CustomWriter) Write(p []byte) (int, error) {
	formatted := fmt.Sprintf("[%s] %s\n", custom.Prefix, string(p))

	return fmt.Fprint(os.Stdout, formatted)
}

func NewColorWriter(col int) *ColorWriter {
	return &ColorWriter{Color: col}
}

func NewCustomWriter(prefix string) *CustomWriter {
	return &CustomWriter{Prefix: prefix}
}

func main() {
	custom := NewCustomWriter("test")

	col := NewColorWriter(3)
	custom.Write([]byte("zhopa"))
	col.Write([]byte("zhopa"))

}
