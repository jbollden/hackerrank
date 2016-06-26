// Day 17: More Exceptions (https://www.hackerrank.com/challenges/30-more-exceptions)
// Since Go is not a selectable language this solution has not been tested on HackerRank

package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Calculator struct{}

func (c Calculator) power(n, p float64) (float64, error) {
	if n < 0 || p < 0 {
		return 0, errors.New("Bad String")
	}

	return math.Pow(n, p), nil
}

func getInputStr() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	return strings.TrimSpace(input), err
}

func main() {
	var T int
	fmt.Scan(&T)

	for ; T > 0; T-- {
		input, _ := getInputStr()
		vals := strings.Split(input, " ")
		n, _ := strconv.ParseFloat(vals[0], 64)
		p, _ := strconv.ParseFloat(vals[1], 64)

		calc := Calculator{}
		ans, err := calc.power(n, p)
		if err != nil {
			fmt.Printf("%v\n", err)
		} else {
			fmt.Printf("%v\n", ans)
		}
	}
}
