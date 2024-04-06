package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

    total := 0

	sc := bufio.NewScanner(file)
	for sc.Scan() {

        line := []rune{}
        for _, c := range sc.Text() {
            if unicode.IsDigit(c) {
                line = append(line, c)
            }
        }

        l := line[0] - '0'
        r := line[len(line) - 1] - '0'

        total += int(10*l + r)
	}

    fmt.Println(total)
}
