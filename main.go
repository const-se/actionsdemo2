package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fmt.Println(greeting(os.Args[1]))
	}
}

func greeting(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
