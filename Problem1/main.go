package problem1

import "fmt"

func main() {
	fmt.Println(twosum([]int{10, 15, 3, 7}, 17))
	fmt.Println(twosum([]int{10, 15, 3, 7}, -15))
	fmt.Println(twosum([]int{0, -15, 3, 7}, -15))
}

// twosum uses a map to store the difference of the current index and the k value
// this lets us store all of the values we pass through, along with the counterpart
// required to add to k, and check if we can return true as we iterate through `nums`.
func twosum(nums []int, k int) bool {
	diffmap := make(map[int]int)
	// loop through our numbers
	for _, num := range nums {
		// if the current value is already an index of the map, there is a pair that adds to k
		if _, ok := diffmap[num]; ok {
			return true
		}
		// if the current value is not an index of our map, add k-num as an index.
		// that way if `k-num` is a value in our list that we will find later, we will
		// be able to look it up in O(n) time and return true
		diffmap[k-num] = num
	}
	return false
}
