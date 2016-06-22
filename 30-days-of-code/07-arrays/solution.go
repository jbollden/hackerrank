// Day 7: Arrays (https://www.hackerrank.com/challenges/30-arrays)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)

	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	A := strings.Split(s, " ")

	var res string
	for i := N - 1; i >= 0; i-- {
		res = res + string(A[i]) + " "
	}
	fmt.Printf("%v\n", res)
}
