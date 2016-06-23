// Day 13: Abstract Classes (https://www.hackerrank.com/challenges/30-abstract-classes)
// Since Go is not a selectable language this solution has not been tested on HackerRank

package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

type IBook interface {
	display()
}

type Book struct {
	title string
	author string
}

type MyBook struct {
	Book
	price int
}

func (m MyBook) display() {
	fmt.Printf("Title: %v\n", m.title)
	fmt.Printf("Author: %v\n", m.author)
	fmt.Printf("Price: %v\n", m.price)
}

func main() {
	var title, author string
	var price int

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < 3; i++ {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch i {
		case 0:
			title = input
		case 1:
			author = input
		case 2:
			price, _ = strconv.Atoi(input)
		}
	}

	var novel IBook = MyBook{Book{title, author}, price}
	novel.display()
}
