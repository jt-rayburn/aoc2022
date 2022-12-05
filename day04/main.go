package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

	// Part A
	aScore := 0
	for _, i := range lines {
		s := strings.Split(i, ",") // [2-4 6-8]
		a := strings.Split(s[0], "-")
		aMin, _ := strconv.Atoi(a[0])
		aMax, _ := strconv.Atoi(a[1])

		b := strings.Split(s[1], "-")
		bMin, _ := strconv.Atoi(b[0])
		bMax, _ := strconv.Atoi(b[1])
		// fmt.Println(aMin, aMax, bMin, bMax)

		if bMin >= aMin && bMax <= aMax {
			aScore += 1
			fmt.Println(aMin, aMax, bMin, bMax)
		} else if aMin >= bMin && aMax <= bMax {
			aScore += 1
			fmt.Println(aMin, aMax, bMin, bMax)
		}

	}
	fmt.Println(aScore)

	// Part B
	bScore := 0
	for _, i := range lines {
		s := strings.Split(i, ",") // [2-4 6-8]
		a := strings.Split(s[0], "-")
		aMin, _ := strconv.Atoi(a[0])
		aMax, _ := strconv.Atoi(a[1])

		b := strings.Split(s[1], "-")
		bMin, _ := strconv.Atoi(b[0])
		bMax, _ := strconv.Atoi(b[1])

		if bMin >= aMin && bMin <= aMax || bMax <= aMax && bMax >= aMin {
			bScore += 1
			fmt.Println(aMin, aMax, bMin, bMax)
		} else if aMin >= bMin && aMin <= bMax || aMax <= bMax && aMax >= bMin {
			bScore += 1
			fmt.Println(aMin, aMax, bMin, bMax)
		}
	}
	fmt.Println(bScore)
}
