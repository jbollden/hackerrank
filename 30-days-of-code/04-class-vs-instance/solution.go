// Day 4: Class vs. Instance (https://www.hackerrank.com/challenges/30-class-vs-instance)

package main

import "fmt"

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	if initialAge < 0 {
		fmt.Println("Age is not valid, setting age to 0.")
		p.age = 0
	} else {
		p.age = initialAge
	}

	return p
}

func (p person) amIOld() {
	res := "You are old."
	if p.age < 13 {
		res = "You are young."
	} else if p.age < 18 {
		res = "You are a teenager."
	}
	fmt.Printf("%v\n", res)
}

func (p person) yearPasses() person {
	p.age++
	return p
}

func main() {
	var T, age int

	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&age)
		p := person{age: age}
		p = p.NewPerson(age)
		p.amIOld()
		for j := 0; j < 3; j++ {
			p = p.yearPasses()
		}
		p.amIOld()
		fmt.Println()
	}
}
