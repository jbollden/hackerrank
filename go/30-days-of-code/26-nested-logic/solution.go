// Day 26: Nested Logic (https://www.hackerrank.com/challenges/30-nested-logic)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	DAY   = 0
	MONTH = 1
	YEAR  = 2
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	actually, _ := reader.ReadString('\n')
	expected, _ := reader.ReadString('\n')

	var actuallyDate, expectedDate [3]int
	for i, v := range strings.Split(actually, " ") {
		actuallyDate[i], _ = strconv.Atoi(strings.TrimSpace(v))
	}
	for i, v := range strings.Split(expected, " ") {
		expectedDate[i], _ = strconv.Atoi(strings.TrimSpace(v))
	}

	fine := 0
	if actuallyDate[YEAR] > expectedDate[YEAR] {
		fine = 10000
	} else if actuallyDate[YEAR] == expectedDate[YEAR] {
		if actuallyDate[MONTH] > expectedDate[MONTH] {
			fine = 500 * (actuallyDate[MONTH] - expectedDate[MONTH])
		} else if actuallyDate[DAY] > expectedDate[DAY] {
			fine = 15 * (actuallyDate[DAY] - expectedDate[DAY])
		}
	}

	fmt.Printf("%v\n", fine)
}
