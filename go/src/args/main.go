package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)

	for _, arg := range os.Args {
		fmt.Println(arg)
	}
}
