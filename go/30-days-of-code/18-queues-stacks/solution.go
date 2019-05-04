// Day 18: Queues and Stacks
// Since Go is not a selectable language this solution has not been tested on HackerRank

package main

import "fmt"

type Solution struct {
	stack []byte
	queue []byte
}

func (s *Solution) pushCharacter(ch byte) {
	s.stack = append(s.stack, ch)
}
func (s *Solution) enqueueCharacter(ch byte) {
	s.queue = append(s.queue, ch)
}
func (s *Solution) popCharacter() byte {
	ch := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return ch
}
func (s *Solution) dequeueCharacter() byte {
	ch := s.queue[0]
	s.queue = s.queue[1:]
	return ch
}

func main() {
	var s string
	fmt.Scan(&s)

	obj := Solution{}

	length := len(s)
	isPalindrome := true

	for i := 0; i < length; i++ {
		obj.pushCharacter(s[i])
		obj.enqueueCharacter(s[i])
	}

	for i := 0; i < length/2; i++ {
		if obj.popCharacter() != obj.dequeueCharacter() {
			isPalindrome = false
			break
		}
	}

	if isPalindrome {
		fmt.Printf("The word, %s, is a palindrome.\n", s)
	} else {
		fmt.Printf("The word, %s, is not a palindrome.\n", s)
	}
}
