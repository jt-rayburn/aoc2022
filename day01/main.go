package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
func main() {

	f, err := os.Open("input.txt")
	check(err)

	defer f.Close()
	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Part 1
	var highest int
	var curr int
	for _, v := range lines {
		if v != "" {
			num, err := strconv.Atoi(v)
			check(err)
			curr = curr + num
		} else {
			if curr > highest {
				highest = curr
			}
			curr = 0
		}
	}
	fmt.Println("Part 1: ", highest)

	// Part 2
	var a int
	var b int
	var c int
	var cur int
	for _, v := range lines {
		if v != "" {
			num, err := strconv.Atoi(v)
			check(err)
			cur = cur + num
		} else {
			switch {
			case cur > a:
				a = cur
			case cur > b:
				b = cur
			case cur > c:
				c = cur
			}
			cur = 0
		}
	}
	total := a + b + c
	fmt.Println("Part 2: ", total)
}
