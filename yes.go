package main

import (
	"fmt"
)

func main() {
	expletive := "y\n"
	bb := []byte(expletive)

	for {
		fmt.Print(bb)
		// os.Stdout.Write(bb)
	}
}
