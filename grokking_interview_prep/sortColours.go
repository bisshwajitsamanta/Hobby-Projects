package main

import "fmt"

func sortColours(colors []int) []int {
	left, middle, right := 0, 0, len(colors)-1
	for middle <= right {
		if colors[middle] == 0 {
			colors[left], colors[middle] = colors[middle], colors[left]
			left++
			middle++
		} else if colors[middle] == 1 {
			middle++
		} else {
			colors[middle], colors[right] = colors[right], colors[middle]
			right--
		}
	}
	return colors
}

func main() {
	fmt.Println(sortColours([]int{1, 0, 2, 1, 0, 2, 1, 0}))
}
