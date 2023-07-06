package main

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	input, err := readInput()
	if err != nil {
		fmt.Println(err)
		return
	}

	var result int32

	if input == 0 {
		result = 0
	} else {
		result, err = reverseInt32(input)
		if err != nil {
			fmt.Println("error - more then int32: ", err)
			return
		}
	}

	fmt.Println("\nresult: ", result)
}

func reverseInt32(i int32) (int32, error) {
	var (
		res int32
		err error
	)

	for i != 0 {
		ch := i / 10
		ost := i % 10

		i = ch

		res, err = addToRes(res, ost)
		if err != nil {
			return 0, err
		}

		if ch == 0 {
			return res, nil
		}
	}

	return res, nil
}

func addToRes(cur int32, new int32) (int32, error) {
	if cur > math.MaxInt32/10 || cur < math.MinInt32/10 {
		return 0, fmt.Errorf("result exceeds int32 range")
	}

	if (cur == math.MaxInt32/10 && new > math.MaxInt32%10) || (cur == math.MinInt32/10 && new < math.MinInt32%10) {
		return 0, fmt.Errorf("result exceeds int32 range")
	}

	cur = cur*10 + new

	return cur, nil
}

func readInput() (int32, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter int32:")

	scanner.Scan()

	input := scanner.Text()

	i, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	i32 := int32(i)

	if i32 < math.MinInt32 || i32 > math.MaxInt32 {
		return 0, errors.New("not int 32")
	}

	return i32, nil
}
