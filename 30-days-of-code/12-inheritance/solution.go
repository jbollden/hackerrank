// Day 12: Inheritance (https://www.hackerrank.com/challenges/30-inheritance)
// Since Go is not a selectable language this solution has not been tested on HackerRank

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
	id        int
}

func (p Person) printPerson() {
	fmt.Printf("Name: %v, %v\nID: %v\n", p.lastName, p.firstName, p.id)
}

type Student struct {
	Person
	testScores []int
}

func (s Student) calculate() string {
	var sum int
	var avg float64
	var letter string

	for _, v := range s.testScores {
		sum += v
	}

	avg = float64(sum) / float64(len(s.testScores))

	switch {
	case 90 <= avg && avg <= 100:
		letter = "O"
	case 80 <= avg && avg < 90:
		letter = "E"
	case 70 <= avg && avg < 80:
		letter = "A"
	case 55 <= avg && avg < 70:
		letter = "P"
	case 40 <= avg && avg < 55:
		letter = "D"
	default:
		letter = "T"
	}

	return letter
}

func main() {
	var person []string
	var numScores int
	var scores []int

	var firstName, lastName string
	var id int

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 3; i++ {
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch i {
		case 0:
			person = strings.Split(input, " ")
			firstName = person[0]
			lastName = person[1]
			id, _ = strconv.Atoi(person[2])
		case 1:
			numScores, _ = strconv.Atoi(input)
		case 2:
			for j, v := range strings.Split(input, " ") {
				if j == numScores {
					break
				}
				score, _ := strconv.Atoi(v)
				scores = append(scores, score)
			}
		}
	}

	student := Student{Person{firstName, lastName, id}, scores}
	student.printPerson()
	fmt.Printf("Grade: %v\n", student.calculate())
}
