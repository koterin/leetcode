package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := readInput()

	pal := findPalindrom(input)

	fmt.Println(pal)
}

func findPalindrom(input string) string {
	pals := make([]string, 0)

	if isPalindrom(input) && len(input) > 1 {
		pals = append(pals, input)
	}

	for i := 0; i < len(input); i++ {
		for j := i + 1; j <= len(input); j++ {
			f := input[i:j]

			if isPalindrom(f) && len(f) > 1 {
				pals = append(pals, f)
			}
		}
	}

	if len(pals) < 1 {
		return input[:1]
	}

	max := pals[0]

	for _, p := range pals {
		if len(p) > len(max) {
			max = p
		}
	}

	return max
}

func isPalindrom(candidate string) bool {
	for i := 0; i < len(candidate); i++ {
		if candidate[i] != candidate[len(candidate)-1-i] {
			return false
		}
	}

	return true
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter string: ")

	scanner.Scan()

	return scanner.Text()
}
