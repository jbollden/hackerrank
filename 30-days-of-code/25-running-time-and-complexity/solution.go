// Day 25: Running Time and Complexity (https://www.hackerrank.com/challenges/30-running-time-and-complexity)

package main

import "fmt"

func is_prime(n int) bool {
	if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if n % 2 == 0 || n % 3 == 0 {
		return false
	}
	for i := 5; i * i <= n; i++ {
		if n % i == 0 || n % (i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	var T, n int

	fmt.Scan(&T)

	for ; T > 0; T-- {
		fmt.Scan(&n)

		if is_prime(n) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
		}
	}
}
