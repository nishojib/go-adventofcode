package day01

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Part1SolveProblem() {
	data, err := os.ReadFile("2015/day01/input.txt")
	check(err)

	counter := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '(' {
			counter++
		} else if data[i] == ')' {
			counter--
		}
	}

	fmt.Println(counter)
}

func Part2SolveProblem() {
	data, err := os.ReadFile("2015/day01/input.txt")
	check(err)

	counter := 0
	basementStep := 0
	for i := 0; i < len(data); i++ {
		if data[i] == '(' {
			counter++
		} else if data[i] == ')' {
			counter--
		}

		if counter < 0 {
			basementStep = i + 1
			break
		}
	}

	fmt.Println(basementStep)
}
