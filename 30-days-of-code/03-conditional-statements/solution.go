// Day 3: Intro to Conditional Statements (https://www.hackerrank.com/challenges/30-conditional-statements)

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	res := "Not Weird"

	if n%2 != 0 || (6 <= n && n <= 20) {
		res = "Weird"
	}

	fmt.Printf("%v\n", res)
}
