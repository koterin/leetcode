package main

import (
	"fmt"
	"strconv"
)

func main() {
	mapping := make(map[int][]string, 8)
	mapping[2] = []string{"a", "b", "c"}
	mapping[3] = []string{"d", "e", "f"}
	mapping[4] = []string{"g", "h", "i"}
	mapping[5] = []string{"j", "k", "l"}
	mapping[6] = []string{"m", "n", "o"}
	mapping[7] = []string{"p", "q", "r", "s"}
	mapping[8] = []string{"t", "u", "v"}
	mapping[9] = []string{"w", "x", "y", "z"}

	input := "2567"
	digits, err := extractDigits(input)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	res := findCombinations(mapping, digits)

	fmt.Println("combinations: ", res)
}

func extractDigits(d string) ([]int, error) {
	digits := make([]int, 0)

	if len(d) < 1 {
		return []int{}, nil
	}

	for _, n := range d {
		new, err := strconv.Atoi(string(n))
		if err != nil {
			return nil, err
		}

		digits = append(digits, new)
	}

	return digits, nil
}

func findCombinations(m map[int][]string, d []int) []string {
	combinations := make([]string, 0)

	if len(d) < 1 {
		return []string{}
	}

	for _, n := range m[d[0]] {
		combinations = append(combinations, n)
	}

	for i := 1; i < len(d); i++ {
		combinations = addCombinations(combinations, m[d[i]])
	}

	return combinations
}

func addCombinations(c []string, m []string) []string {
	newComb := make([]string, 0)

	for i := 0; i < len(c); i++ {
		for _, n := range m {
			newComb = append(newComb, c[i]+n)
		}
	}

	return newComb
}
