package main

import (
	"fmt"
	"os"
)

func main() {
	expletive := "y"
	if len(os.Args) != 1 {
		expletive = os.Args[1]
	}
	expletive += "\n"

	bb := []byte(expletive)

	for {
		fmt.Print(bb)
	}
}
