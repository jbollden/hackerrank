// Day 14: Scope (https://www.hackerrank.com/challenges/30-scope)
// Since Go is not a selectable language this solution has not been tested on HackerRank

package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"strconv"
)

type Difference struct {
	elements []int
	MaximumDifference int
}
func (d *Difference) computeDifference() {
	for i, v := range d.elements {
		for j := i; j < len(d.elements); j++ {
			diff := d.elements[j] - v
			if diff > d.MaximumDifference {
				d.MaximumDifference = diff
			}
		}
	}
}

func getInputStr() (string, error) {
	reader := bufio.NewReader(os.Stdin)
    input, err := reader.ReadString('\n')
    return strings.TrimSpace(input), err
}

func main() {
	var N int
	var elements []int

	for i := 0; i < 2; i++ {
		input, _ := getInputStr()
		switch i {
		case 0:
			N, _ = strconv.Atoi(input)
		case 1:
			vals := strings.Split(input, " ")
			for i := 0; i < N; i++ {
				v := 0
				if i < len(vals) {
					v, _ = strconv.Atoi(vals[i])
				}
				elements = append(elements, v)
			}
		}
	}

	d := Difference{elements: elements}
	d.computeDifference()
	fmt.Printf("%v\n", d.MaximumDifference)
}
