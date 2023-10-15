package main

import (
	"fmt"
	"io"
	"os"
)

// type logWriter struct{}

func main() {
	// fmt.Println(os.Args[1])

	file, err := os.Open("file.txt") // For read access.

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// lw := logWriter{}

	io.Copy(os.Stdout, file)
}

// func (logWriter) Write(bs []byte) (int, error) {
// 	fmt.Println(string(bs))
// 	// fmt.Println("Just wrote this many bytes:", len(bs))
// 	return len(bs), nil
// }
