// Day 0: Hello, World. (https://www.hackerrank.com/challenges/30-hello-world)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputString, _ := reader.ReadString('\n')
	inputString = strings.TrimSpace(inputString)

	fmt.Println("Hello, World.")
	fmt.Printf("%v\n", inputString)
}
