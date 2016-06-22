// Day 1: Data Types (https://www.hackerrank.com/challenges/30-data-types)
// Since Go is not a selectable language this solution has not been tested on HackerRank

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	I := 4
	D := 4.0
	S := "HackerRank "

	var i int
	var d float64
	var s string

	fmt.Scan(&i)
	fmt.Scan(&d)

	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	s = strings.TrimSpace(s)

	fmt.Printf("%v\n", I+i)
	fmt.Printf("%01.1f\n", D+d)
	fmt.Printf("%v\n", S+s)
}
