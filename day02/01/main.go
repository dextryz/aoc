package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/scanner"
)

// [ red, green, blue ]

func isPossible(hands [][]int, cubes []int) bool {

	for _, hand := range hands {
        for i := 0; i < len(hand); i++ {
            if hand[i] > cubes[i] {
                return false
            }
        }
	}

	return true
}

// 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func tokenize(line string) [][]int {

	colLen := len(strings.Split(line, ";"))

	res := make([][]int, 0)
	for i := 0; i < colLen; i++ {
		res = append(res, make([]int, 3))
	}

	var s scanner.Scanner
	s.Init(strings.NewReader(line))

	// treat leading '%' as part of an identifier
	s.IsIdentRune = func(ch rune, i int) bool {
		return ch != ',' && ch != ';'
	}

	row := 0

	r := make([]int, 3)

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {

		switch tok {
		case ',':
		case ';':
			row++
			r = make([]int, 3)
		default:

			hand := strings.Split(s.TokenText(), " ")

			switch hand[1] {
			case "red":
				r[0], _ = strconv.Atoi(hand[0])
			case "green":
				r[1], _ = strconv.Atoi(hand[0])
			case "blue":
				r[2], _ = strconv.Atoi(hand[0])
			}

			res[row] = r
		}
	}

	return res
}

func main() {

	file, _ := os.Open("input")

	sc := bufio.NewScanner(file)

	sum := 0
	count := 1

	for sc.Scan() {

		prefix := fmt.Sprintf("Game %d: ", count)
		line := strings.TrimPrefix(sc.Text(), prefix)

		tokens := tokenize(line)

		if isPossible(tokens, []int{12, 13, 14}) {
			sum += count
		}

		count++
	}

	fmt.Println(sum)
}
