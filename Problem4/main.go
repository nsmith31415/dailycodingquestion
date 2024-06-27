package problem4

import (
	"fmt"
)

func main() {
	fmt.Println(sort([]string{"G", "B", "R", "R", "B", "R", "B"}))
	fmt.Println(sort([]string{"B", "B", "G", "G", "R", "R", "R"}))
	fmt.Println(sort([]string{"G", "B", "R", "R", "B", "R", "G", "B", "R", "R", "B", "R", "G", "B", "R", "R", "B", "R", "G"}))

}

// the first solution that comes to mind:
//
// A greedy solution. We loop through the array 2 times. The first time we loop through,
// we're looking for Bs. When we find a B, we set it to be the last element in the array
// that is not already a B. The second loop through, we repeat for G, asetting it to the
// last element that is not a B or a G. When this loop is done, all of the Rs should be
// the first elements in the array
func sort(arr []string) []string {
	// since we're sorting and filling in from the back, count will track the
	// number of elements we have correct. We don't want to disrupt anything past
	// the index of len(arr) - (count + 1)
	count := 0
	// loop through the array looking for any letter B that isn't yet sorrted
	for i, l := range arr {
		// if we find an unsorted B, we'll then find the last unsorted index
		if len(arr)-i > count && l == "B" {
			// swapped will let us break out of the loop quicker so we don't go through all of arr here
			swapped := false
			// backIndex gives us the last index that we don't know to be sorted, which we'll then check
			backIndex := len(arr) - (count + 1)
			count++
			for backIndex > i && !swapped {
				if arr[backIndex] != "B" { // if the backIndex isn't sorted, we can swap it
					swapped = true
					arr[i], arr[backIndex] = arr[backIndex], arr[i]
				} else { // if the backIndex is sorted, then we can increment back toward the front of arr
					backIndex--
					// we'll also add to count because we've just iterated past an element that happened to be sorted already
					count++
				}
			}
		}
	}

	// Do the same as agove, but for the Gs
	for i, l := range arr {
		if len(arr)-i > count && l == "G" {
			swapped := false
			backIndex := len(arr) - (count + 1)
			count++
			for backIndex > i && !swapped {
				// Because all of the Bs are now sorted, and we won't bother swapping a G with itself,
				// we only have to look for an R to swap the G with.
				if arr[backIndex] == "R" {
					swapped = true
					arr[i], arr[backIndex] = arr[backIndex], arr[i]
				} else {
					backIndex--
					count++
				}
			}
		}
	}
	return arr
}
