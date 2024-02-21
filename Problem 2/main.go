package main

import "fmt"

func main() {
	fmt.Println(arrayproduct([]int{10, 15, 3, 7}))
	fmt.Println(arrayproduct([]int{2, 4, 6, 8}))
	fmt.Println(arrayproductnodivision([]int{10, 15, 3, 7}))
	fmt.Println(arrayproductnodivision([]int{2, 4, 6, 8}))
}

// The easiest solution to this problem is to take the product of the
// array of integers, and then create a new array by dividing the product
// by the value at index `i` of the original array. This solution is shown
// below
func arrayproduct(nums []int) []int {
	product := 1

	//first we create the product
	for _, val := range nums {
		product = product * val
	}

	//now we create and assign the new array
	final := make([]int, len(nums))
	for i, _ := range final {
		final[i] = product / nums[i]
	}
	return final
}

// For the follow-up, we are asked to find a solution that does not use division:
// This solution creates two arrays from the original before calculating the solution.
// The first array is an array such that at index `i`, the value is the product of all
// elements to the right of index `i` in the original array.
// The second array is an array such that at index `i`, the value is the product of all
// elements to the left of index `i` in the original array.
// We can then calculate the final solution by finding the product of the values of this
// right array and left array for a given index `i`.
func arrayproductnodivision(nums []int) []int {
	// set the first value of the right array
	rightArray := make([]int, len(nums))
	rightArray[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		rightArray[i] = nums[i+1] * rightArray[i+1]
	}

	// set the first value of the left array
	leftArray := make([]int, len(nums))
	leftArray[0] = 1
	for i := 1; i < len(nums); i++ {
		leftArray[i] = nums[i-1] * leftArray[i-1]
	}

	final := make([]int, len(nums))
	for i, _ := range final {
		final[i] = rightArray[i] * leftArray[i]
	}
	return final
}
