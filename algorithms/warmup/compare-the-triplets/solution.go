// Algorithms / Warmup / Compare the Triplets (https://www.hackerrank.com/challenges/compare-the-triplets)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var a []int
	for _, v := range strings.Split(input, " ") {
		val, _ := strconv.Atoi(v)
		a = append(a, val)
	}

	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var b []int
	for _, v := range strings.Split(input, " ") {
		val, _ := strconv.Atoi(v)
		b = append(b, val)
	}

	var A, B int
	for i := 0; i < 3; i++ {
		if a[i] < b[i] {
			B++
		} else if a[i] > b[i] {
			A++
		}
	}
	fmt.Println(A, B)
}
