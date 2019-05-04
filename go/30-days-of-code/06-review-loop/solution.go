// Day 6: Let's Review (https://www.hackerrank.com/challenges/30-review-loop)

package main

import "fmt"

func main() {
	var T int
	var S, odd, even string

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&S)
		for j, v := range S {
			if j%2 == 0 {
				even = even + string(v)
			} else {
				odd = odd + string(v)
			}
		}
		fmt.Printf("%v %v\n", even, odd)
		even, odd = "", ""
	}
}
