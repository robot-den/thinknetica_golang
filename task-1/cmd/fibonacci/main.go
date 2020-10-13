package main

import (
	"../../pkg/fibonacci"
	"flag"
	"fmt"
)

var nFlag = flag.Int("n", 1, "position of required number (must be in range 1..20)")

func main() {
	flag.Parse()
	fmt.Println(fibonacci.At(*nFlag - 1))
}
