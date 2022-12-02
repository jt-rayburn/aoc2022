package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	// Key
	// A, X = Rock  = 1pt
	// B, Y = Paper = 2pt
	// C, Z = Scis  = 3pt
	// Win  = 6pt
	// Draw = 3pt
	// Lose = 0pt

	// Part 1
	var score int
	for i := 0; i < len(lines); i++ {
		a := strings.Split(lines[i], " ")
		switch {
		case a[0] == "A":
			switch {
			case a[1] == "X":
				score += 4 // draw, rock
			case a[1] == "Y":
				score += 8 // win, pape
			case a[1] == "Z":
				score += 3 // lose, scis
			}
		case a[0] == "B":
			switch {
			case a[1] == "X":
				score += 1 // lose, rock
			case a[1] == "Y":
				score += 5 // draw, pape
			case a[1] == "Z":
				score += 9 // win, scis
			}
		case a[0] == "C":
			switch {
			case a[1] == "X":
				score += 7 // win, rock
			case a[1] == "Y":
				score += 2 // lose, pape
			case a[1] == "Z":
				score += 6 // draw, scis
			}
		}
	}
	fmt.Println(score)

	// Part 2

	// Key
	// X Lose, Y Draw, Z Win
	var newScore int
	for i := 0; i < len(lines); i++ {
		a := strings.Split(lines[i], " ")
		switch {
		case a[0] == "A":
			switch {
			case a[1] == "X":
				newScore += 3 // LOSE, SCIS
			case a[1] == "Y":
				newScore += 4 // DRAW, ROCK
			case a[1] == "Z":
				newScore += 8 // WIN, PAPER
			}
		case a[0] == "B":
			switch {
			case a[1] == "X":
				newScore += 1 // LOSE, ROCK
			case a[1] == "Y":
				newScore += 5 // DRAW, PAPE
			case a[1] == "Z":
				newScore += 9 // WIN, SCIS
			}
		case a[0] == "C":
			switch {
			case a[1] == "X":
				newScore += 2 // LOSE, PAPER
			case a[1] == "Y":
				newScore += 6 // DRAW, SCIS
			case a[1] == "Z":
				newScore += 7 // WIN, ROCK
			}
		}
	}
	fmt.Println(newScore)
}
