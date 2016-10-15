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

	fmt.Scan(&N)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var a []int
	for _, v := range strings.Split(input, " ") {
		val, _ := strconv.Atoi(v)
		a = append(a, val)
	}

	var sum int
	for i := 0; i < N; i++ {
		sum += a[i]
	}
	fmt.Println(sum)
}
