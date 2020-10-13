// Package fibonacci implements functions to work with fibonacci sequence
package fibonacci

import (
	"log"
)

// At calculates fibonacci value at specified index
func At(n int) int {
	if n < 0 || n > 19 {
		log.Fatal("position value is invalid")
	} else if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	values := make([]int, n+1, n+1)
	values[1] = 1

	for i := 2; i <= n; i++ {
		values[i] = values[i-2] + values[i-1]
	}

	return values[n]
}

// Recursive At() variant
// func At(n int) int {
//   if n < 0 || n > 19 {
// 		log.Fatal("position value is invalid")
// 	} else if n == 0 {
// 		return 0
// 	} else if n == 1 {
// 		return 1
// 	}
//   return At(n-1) + At(n-2)
// }
