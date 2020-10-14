// Package fibonacci implements functions to work with fibonacci sequence
package fibonacci

// At calculates fibonacci value at specified index
func At(n int) (int, error) {
	if n == 0 {
		return 0, nil
	} else if n == 1 {
		return 1, nil
	}

	first, second := 0, 1

	for i := 2; i <= n; i++ {
		first, second = second, first + second
	}

	return second, nil
}

// Recursive At() variant
// func At(n int) int {
//   if n < 0 || n > 20 {
// 		log.Fatal("position value is invalid")
// 	} else if n == 0 {
// 		return 0
// 	} else if n == 1 {
// 		return 1
// 	}
//   return At(n-1) + At(n-2)
// }
