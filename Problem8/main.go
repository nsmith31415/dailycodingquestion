package main

import "fmt"

func main() {
	fmt.Println(findSmallestSet([][]int{{0, 3}, {2, 6}, {3, 4}, {6, 9}}))
	fmt.Println(findSmallestSet([][]int{{0, 7}, {2, 21}, {15, 20}, {6, 9}}))
}

// theoretically speaking, if an interval is defined by [x,y], the
// interval of [smallest y, largest x] should create a set that
// fits all intervals.

// assumptions: this is a valid input of a set of intervals, with at
// least one interval, and each interval has exactly 2 integers.
func findSmallestSet(intervals [][]int) []int {
	smallestY := intervals[0][1]
	largestX := intervals[0][0]
	for i := range intervals {
		if smallestY > intervals[i][1] {
			smallestY = intervals[i][1]
		}
		if largestX < intervals[i][0] {
			largestX = intervals[i][0]
		}
	}

	// if the smallest Y is greater than the largestX at the end of
	// the loop, then all checked intervals should be representable
	// by a single digit, we will choose the smaller of the two.
	if smallestY > largestX {
		smallestY = largestX
	}
	return []int{smallestY, largestX}
}
