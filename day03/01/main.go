package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// FAILED
// 6933835 (new line break num)
// 8046507 (new line does not break num)
// 8909911
// 518930
// 1634187

// TRY

func dimensions(file *os.File) (int, int) {
	sc := bufio.NewScanner(file)
	rowCnt, colCnt := 0, 0
	for sc.Scan() {
		colCnt = len(sc.Text())
		rowCnt++
	}
	return rowCnt + 2, colCnt + 2
}

func printMatrix(mat [][]rune, n, m int) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%c", mat[i][j])
		}
		fmt.Println("")
	}
}

func main() {

	file, _ := os.Open("input")

	rows, cols := dimensions(file)

	mat := make([][]rune, 0)
	for i := 0; i < rows; i++ {
		mat = append(mat, make([]rune, cols))
	}

	fmt.Printf("%d %d\n", rows, cols)

	// 1. Poplate the matrix boundaries

	for i := 0; i < cols; i++ {
		mat[0][i] = '.'
		mat[rows-1][i] = '.'
	}
	for i := 0; i < rows; i++ {
		mat[i][0] = '.'
		mat[i][cols-1] = '.'
	}

	// 2. Populate the matrix with the input data

	// Reset file pointer to reread buffer data.
	_, _ = file.Seek(0, io.SeekStart)

	sc := bufio.NewScanner(file)
	row := 1
	for sc.Scan() {
		for col, v := range []rune(sc.Text()) {
			mat[row][col+1] = v
		}
		row++
	}

	// 3. Create boundary checking loop

	boundaries := [][]int{
		{-1, -1}, // top-left
		{-1, 0},  // top-center
		{-1, 1},  // top-right
		{0, 1},   // right
		{1, 1},   // bot-right
		{1, 0},   // bot-center
		{1, -1},  // bot-left
		{0, -1},  // left
	}

	//printMatrix(mat, rows, cols)

	sum := 0
	num := 0
	keep := false

	for i := 1; i < rows; i++ {
		for j := 1; j < cols; j++ {

			r := mat[i][j]

			if unicode.IsDigit(r) {

				num = 10*num + int(r-'0')

                // Carful, this boundary search will never happen for non-digits.
				for _, b := range boundaries {
					di := i + b[0]
					dj := j + b[1]
					s := mat[di][dj]
					if s != '.' && !unicode.IsDigit(s) {
						keep = true
					}
				}
			} else {
				if keep {
					sum += num
					keep = false
				}
				num = 0
			}
		}
	}

	fmt.Println(sum)
}
