// Algorithms / Warmup / Staircase (https://www.hackerrank.com/challenges/staircase)

package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		fmt.Print(strings.Repeat(" ", n-i))
		fmt.Println(strings.Repeat("#", i))
	}
}
