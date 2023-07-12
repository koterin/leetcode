package main

import "fmt"

func main() {
	var n int

	fmt.Scanln(&n)

	pairs := makeCombinations(n)
	output := formPairs(pairs, n)

	fmt.Println(output)
}

func formPairs(b []string, n int) []string {
	output := []string{}

	for _, v := range b {
		discard := false
		openCount := 0
		closeCount := 0

		for i := 0; i < len(v); i++ {
			if string(v[i]) == "(" {
				openCount++
				continue
			}

			if (string(v[i]) == ")") && (openCount <= closeCount || openCount < 1) {
				discard = true
				break
			}

			closeCount++
		}

		if (!discard) && (openCount == n) && (closeCount == n) {
			output = append(output, v)
		}
	}

	return output
}

func makeCombinations(n int) []string {
	c := []string{}

	c = append(c, "(")
	c = append(c, ")")

	for i := 0; i < n*2-1; i++ {
		nn := len(c)
		for j := 0; j < nn; j++ {
			cur := c[j]
			c[j] = cur + "("
			c = append(c, cur+")")
		}
	}

	return c
}
