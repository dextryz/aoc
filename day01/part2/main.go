package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var letters = []string{
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func main() {

	file, _ := os.Open("input")

	sc := bufio.NewScanner(file)

	sum := 0

	for sc.Scan() {

		line := make([]int, len(sc.Text()))

		for i, c := range sc.Text() {
			if unicode.IsDigit(c) {
				line[i] = int(c - '0')
			}
		}

		for i, l := range letters {
			id := strings.Index(sc.Text(), l)
			if id != -1 {
				line[id] = i + 1
			}
			id = strings.LastIndex(sc.Text(), l)
			if id != -1 {
				line[id] = i + 1
			}
		}

		i := 0
		for line[i] == 0 {
			i++
		}

		j := len(line) - 1
		for line[j] == 0 {
			j--
		}

		sum += 10*line[i] + line[j]
	}

	fmt.Println(sum)
}
