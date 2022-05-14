package main

import (
	"fmt"
	"os"
)

func main() {
	i := 0
	for {
		i++
		fmt.Println(os.Getpid())
	}
}
