package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

// https://leetcode.com/problems/zigzag-conversion/description/

func main() {
	// input := "PAYPALISHIRING"
	// rows := 4

	input, rows, err := readInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	output := convert(input, rows)

	fmt.Println(output)
}

func convert(s string, numRows int) string {
	res := ""
	rows := make([]string, numRows)

	r := 0          // current row
	next_move := -1 // this variable indicates that we need to move forward or backward

	for i := 0; i < len(s); i++ {
		if r == 0 || r == numRows-1 { // change direction
			next_move = -next_move
		}

		// append a character to the current row
		rows[r] += string(s[i])

		// move to the next row
		r += next_move
	}

	for i := 0; i < len(rows); i++ {
		res += rows[i]
	}

	return res
}

func readInput() (string, int, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a string:")

	scanner.Scan()
	input := scanner.Text()

	fmt.Println("You entered: ", input)

	fmt.Println("Enter number of rows:")

	scanner.Scan()

	rows, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "", 0, errors.New("Invalid input. Please enter a valid integer.")
	}

	fmt.Println("You entered: ", rows)

	return input, rows, nil
}
