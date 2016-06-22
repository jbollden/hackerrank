// Day 10: Binary Numbers (https://www.hackerrank.com/challenges/30-binary-numbers)

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, cnt, max int

	fmt.Scan(&n)
	bin := strconv.FormatInt(int64(n), 2)

	for i := 0; i < len(bin); i++ {
		if string(bin[i]) == "1" {
			cnt++
			if cnt > max {
				max = cnt
			}
		} else {
			cnt = 0
		}
	}

	fmt.Printf("%v\n", max)
}
