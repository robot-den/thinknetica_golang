package main

import (
	"pkg/fibonacci"
	"flag"
	"fmt"
)

var nFlag = flag.Int("n", 0, "index of required number (must be in range 0..20)")

func main() {
	flag.Parse()

	n := *nFlag
	if n < 0 || n > 20 {
		fmt.Println("sorry, we have paws instead of hands and so can only work with indexes in range 0..20")
		return
	}

	result, err := fibonacci.At(n)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
