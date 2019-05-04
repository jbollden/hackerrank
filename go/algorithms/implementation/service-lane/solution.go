// Algorithms / Implementation / Service Lane (https://www.hackerrank.com/challenges/service-lane)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var rows [][]int
	var input string
	var N, T int
	cnt := 2

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < cnt; i++ {
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)

		var row []int
		for _, v := range strings.Split(input, " ") {
			val, _ := strconv.Atoi(v)
			row = append(row, val)
		}
		rows = append(rows, row)

		if i == 0 {
			N = row[0]
			T = row[1]

			cnt = cnt + T
		}
	}

	width := rows[1][:N]
	tests := rows[2 : T+2]

	for i := 0; i < T; i++ {
		res := 5
		for j := tests[i][0]; j <= tests[i][1]; j++ {
			if width[j] < res {
				res = width[j]
			}
		}
		fmt.Println(res)
	}
}
