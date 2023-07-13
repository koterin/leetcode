package main

import "fmt"

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0

	p := findRotatedPoint(nums)

	fmt.Println("pivot: ", p)

	if p == 0 {
		if nums[0] > target {
			fmt.Println(-1)
			return
		}

		fmt.Println(binarySearch(nums, target))
		return
	}

	if nums[0] > target {
		fmt.Println(nums[p:])
		res := binarySearch(nums[p:], target)
		if res == -1 {
			fmt.Println(res)
		} else {
			fmt.Println(res + p)
		}
	} else {
		fmt.Println(nums[:p])
		fmt.Println(binarySearch(nums[:p], target))
	}
}

func findRotatedPoint(array []int) int {
	start := array[0]
	low := 0
	high := len(array) - 1

	if start < array[high] {
		return 0
	}

	for low < high {
		middle := (low + high) / 2

		if array[middle] >= start {
			low = middle + 1
		} else {
			high = middle
		}
	}

	return low
}

func binarySearch(array []int, target int) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		median := (low + high) / 2

		if array[median] < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(array) || array[low] != target {
		return -1
	}

	return low
}
