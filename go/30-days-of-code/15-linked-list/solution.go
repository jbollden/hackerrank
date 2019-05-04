// Day 15: Linked List (https://www.hackerrank.com/challenges/30-linked-list)
// Since Go is not a selectable language this solution has not been tested on HackerRank

/*
NOTE: Compilation of this file will fail with:

    solution.go:11: invalid recursive type Node

However, I solved this day's problem in PHP with the following snippet:

    function insert($head, $data) {
        if ($head == null) {
            $head = new Node($data);
        } else {
            $current = $head;
            while ($current->next != null) {
                $current = $current->next;
            }
            $current->next = new Node($data);
        }
        return $head;
    }
*/

package main

import "fmt"

type Node struct {
	data int
	next Node
}

type Solution struct{}

func (s Solution) insert(head Node, data int) Node {
	if (Node{}) == head {
		head = Node{data: data}
	} else {
		current := head
		for current.next != nil {
			current = current.next
		}
		current.next = Node{data: data}
	}
	return head
}
func (s Solution) display(head Node) {
	start := head
	for (Node{}) != start {
		fmt.Printf("%v ", start.data)
		start = start.next
	}
}

func main() {
	var T, data int

	fmt.Scan(&T)
	head := Node{}
	mylist := Solution{}

	for ; T > 0; T-- {
		fmt.Scan(&data)
		head = mylist.insert(head, data)
	}
	mylist.display(head)
}
