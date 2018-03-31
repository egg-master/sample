package main

import (
	"flag"
	"fmt"
)

var (
	boolflag   = flag.Bool("b", false, "ex -b")
	stringflag = flag.String("s", "first", "ex -s=second")
	intflag    = flag.Int("i", 1, "ex -i=2")
)

func main() {
	flag.Parse()

	fmt.Println(*boolflag)
	fmt.Println(*stringflag)
	fmt.Println(*intflag)
}
