package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func Part1SolveProblem() {
	file, err := os.Open("2015/day02/input.txt")
	check(err)
	defer func(file *os.File) {
		err := file.Close()
		check(err)
	}(file)

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sizesStr := strings.Split(scanner.Text(), "x")
		var sizes []int
		for _, i := range sizesStr {
			size, err := strconv.Atoi(i)
			check(err)
			sizes = append(sizes, size)
		}
		sort.Ints(sizes)

		smallestArea := area(sizes[0], sizes[1])
		sa := surfaceArea(sizes[0], sizes[1], sizes[2])
		sum += sa + smallestArea
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func surfaceArea(l, w, h int) int {
	return 2*l*w + 2*w*h + 2*h*l
}

func area(l, w int) int {
	return l * w
}

func Part2SolveProblem() {
	file, err := os.Open("2015/day02/input.txt")
	check(err)
	defer func(file *os.File) {
		err := file.Close()
		check(err)
	}(file)

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sizesStr := strings.Split(scanner.Text(), "x")
		var sizes []int
		for _, i := range sizesStr {
			size, err := strconv.Atoi(i)
			check(err)
			sizes = append(sizes, size)
		}
		sort.Ints(sizes)

		ribbon := sizes[0]*2 + sizes[1]*2
		bow := sizes[0] * sizes[1] * sizes[2]

		sum += ribbon + bow
	}

	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
