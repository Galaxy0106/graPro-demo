package main

import (
	"fmt"
	"graPro-demo/consumer"
	"os"
)

func main() {
	if os.Args[1] == "-p" {
		NewServer()
	} else if os.Args[1] == "-c" {
		consumer.StartServer()
	} else {
		fmt.Println("Error Args")
	}
}
