// Day 8: Dictionaries and Maps (https://www.hackerrank.com/challenges/30-dictionaries-and-maps)

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N int
	phoneBook := make(map[string]string)

	fmt.Scan(&N)
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < N; i++ {
		input, _ := reader.ReadString('\n')
		entry := strings.Split(strings.TrimSpace(input), " ")
		phoneBook[strings.TrimSpace(entry[0])] = strings.TrimSpace(entry[1])
	}

	var res string
	for {
		query, err := reader.ReadString('\n')
		query = strings.TrimSpace(query)
		if res = phoneBook[query]; res != "" {
			res = query + "=" + res
		} else {
			res = "Not found"
		}
		fmt.Println(res)
		if err != nil {
			return
		}
	}
}
