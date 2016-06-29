// Day 20: Sorting (https://www.hackerrank.com/challenges/30-sorting)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var a []int
	for _, v := range strings.Split(input, " ") {
		val, _ := strconv.Atoi(v)
		a = append(a, val)
	}

	var numberOfTotalSwaps int
	for i := 0; i < n; i++ {
		var numberOfSwaps int
		for j := 0; j < n-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				numberOfSwaps++
			}
		}
		numberOfTotalSwaps += numberOfSwaps
		if numberOfSwaps == 0 {
			break
		}
	}

	fmt.Printf("Array is sorted in %v swaps.\n", numberOfTotalSwaps)
	fmt.Printf("First Element: %v\n", a[0])
	fmt.Printf("Last Element: %v\n", a[len(a)-1])
}
