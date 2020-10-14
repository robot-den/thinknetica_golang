package main

import (
	"pkg/fibonacci"
	"flag"
	"fmt"
)

var nFlag = flag.Int("n", 0, "index of required number (must be in range 0..20)")

func main() {
	flag.Parse()
	result, err := fibonacci.At(*nFlag)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
