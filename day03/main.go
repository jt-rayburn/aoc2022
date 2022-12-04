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

	// Part 1
	scores := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
	}
	var final []string
	total := 0
	for _, i := range lines {
		s := strings.Split(i, "")
		halfS := len(s) / 2
		aHalf := make(map[string]int, halfS)
		bHalf := make(map[string]int, halfS)
		for k, v := range s {
			if k < halfS {
				aHalf[v] = 1
			} else {
				bHalf[v] = 1
			}
		}
		for key, _ := range bHalf {
			if _, ok := aHalf[key]; ok {
				final = append(final, key)

			}
		}
	}
	for _, i := range final {
		if val, ok := scores[i]; ok {
			total += val
		} else {
			i := strings.ToLower(i)
			total += (scores[i]) + 26
		}
	}
	fmt.Println(total)

	// Part 2 - test score 70
	bTotal := 0
	for i := 0; i < len(lines); i = i + 3 {
		a := strings.Split(lines[i], "") // this needs to be a map
		newA := make(map[string]int, len(a))
		for _, j := range a {
			if _, ok := newA[j]; ok {
				newA[j] += 1
			} else {
				newA[j] = 1
			}
		}

		b := strings.Split(lines[i+1], "")
		newB := make(map[string]int, len(b))

		for _, j := range b {
			if _, ok := newB[j]; ok {
				newB[j] += 1
			} else {
				newB[j] = 1
			}
		}

		c := strings.Split(lines[i+2], "")
		newC := make(map[string]int, len(c))

		for _, j := range c {
			if _, ok := newC[j]; ok {
				newC[j] += 1
			} else {
				newC[j] = 1
			}
		}
		for k, _ := range newA {
			if _, ok := newB[k]; ok {
				if _, bok := newC[k]; bok {
					if val, cok := scores[k]; cok {
						bTotal += val
					} else {
						k := strings.ToLower(k)
						bTotal += (scores[k]) + 26
					}
				}
			}
		}

	}
	fmt.Println(bTotal)
}
