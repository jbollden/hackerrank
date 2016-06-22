// Day 9: Recursion (https://www.hackerrank.com/challenges/30-recursion)

package main

import "fmt"

func factorial(n int) int {
	if n <= 1 {
		return n
	}
	return n * factorial(n-1)
}

func main() {
	var N int

	fmt.Scan(&N)

	res := factorial(N)
	fmt.Printf("%v\n", res)
}
