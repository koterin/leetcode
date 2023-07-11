package main

import (
	"fmt"
)

func main() {
	var (
		s1 string
		s2 string
	)

	fmt.Scanln(&s1)
	fmt.Scanln(&s2)

	s1map := make(map[rune]int, len(s1))
	s2map := make(map[rune]int, len(s2))

	if len(s1) != len(s2) {
		fmt.Println(0)
		return
	}

	for _, v := range s1 {
		if _, ok := s1map[v]; ok {
			s1map[v]++
		} else {
			s1map[v] = 1
		}
	}

	for _, v := range s2 {
		if _, ok := s2map[v]; ok {
			s2map[v]++
		} else {
			s2map[v] = 1
		}
	}

	/*
		if ok := reflect.DeepEqual(s1map, s2map); !ok {
			fmt.Println(0)
			return
		}
	*/

	for k1, v1 := range s1map {
		if _, ok := s2map[k1]; !ok {
			fmt.Println(0)
			return
		} else {
			if v1 != s2map[k1] {
				fmt.Println(0)
				return
			}
		}
	}

	fmt.Println(1)

}
