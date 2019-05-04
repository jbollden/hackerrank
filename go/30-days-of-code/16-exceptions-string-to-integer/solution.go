// Day 16: Exceptions - String to Integer (https://www.hackerrank.com/challenges/30-exceptions-string-to-integer)
// Since Go is not a selectable language this solution has not been tested on HackerRank

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var S string
	fmt.Scan(&S)
	val, err := strconv.Atoi(S)
	if err != nil {
		fmt.Println("Bad String")
		return
	}
	fmt.Printf("%v\n", val)
}
