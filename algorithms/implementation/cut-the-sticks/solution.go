// Algorithms / Implementation / Cut the sticks (https://www.hackerrank.com/challenges/cut-the-sticks)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N int
	var sticks []int

	fmt.Scan(&N)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	row := strings.Split(input, " ")
	for i := 0; i < N; i++ {
		val, _ := strconv.Atoi(row[i])
		sticks = append(sticks, val)
	}

	for len(sticks) > 0 {
		var tmp []int
		shortest := getShortest(sticks)

		for i := 0; i < len(sticks); i++ {
			if sticks[i]-shortest > 0 {
				tmp = append(tmp, sticks[i]-shortest)
			}
		}

		fmt.Println(len(sticks))
		sticks = tmp
	}
}

func getShortest(sticks []int) int {
	// Max integer
	shortest := int(^uint(0) >> 1)

	for _, val := range sticks {
		if val < shortest {
			shortest = val
		}
	}
	return shortest
}
