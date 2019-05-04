// Day 11: 2D-Arrays (https://www.hackerrank.com/challenges/30-2d-arrays)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var A [6][6]int
	var maxSum int = -2147483648

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 6; i++ {
		input, _ := reader.ReadString('\n')
		values := strings.Split(strings.TrimSpace(input), " ")

		for j := 0; j < 6; j++ {
			value, _ := strconv.Atoi(values[j])
			A[i][j] = value
		}
	}

	for k := 0; k < len(A)-2; k++ {
		for l := 0; l < len(A[0])-2; l++ {
			sum := A[k][l] + A[k][l+1] + A[k][l+2] + A[k+1][l+1] + A[k+2][l] + A[k+2][l+1] + A[k+2][l+2]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	fmt.Printf("%v\n", maxSum)
}
