package main

import (
	"fmt"
	"io"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name) // Fprintf invokes the Write method of the io.Writer interface
}

func main() {
	Greet(os.Stdout, "Elodie") // os.Stdout implements the io.Writer interface
}
