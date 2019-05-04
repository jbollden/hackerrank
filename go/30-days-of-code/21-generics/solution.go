// Day 21: Generics
// Since Go is not a selectable language this solution has not been tested on HackerRank

// I have found no way to use generics in Go. Converting to interface is the closest I get.

/*
	This is the solution in Java 8:

	public static void printArray(Object[] c) {
		for (Object e : c) {
			System.out.println(e);
		}
	}

*/

package main

import "fmt"

type Solution struct {}
func (s Solution) printArray(arr []interface{}) {
	for _, v := range arr {
		fmt.Printf("%v\n", v)
	}
}

func main() {

	intArray := []int{1, 2, 3}
	stringArray := []string{"Hello", "World"}

	ifIntArray := make([]interface{}, len(intArray))
	ifStrArray := make([]interface{}, len(stringArray))
	for i, v := range intArray {
		ifIntArray[i] = v
	}
	for i, v := range stringArray {
		ifStrArray[i] = v
	}

	s := Solution{}
	s.printArray(ifIntArray)
	s.printArray(ifStrArray)
}
