package main

import (
	"fmt"
)

// binarySearch function performs a binary search on a sorted array
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		// Check if target is present at mid
		if arr[mid] == target {
			return mid
		}

		// If target greater, ignore left half
		if arr[mid] < target {
			left = mid + 1
		} else { // If target is smaller, ignore right half
			right = mid - 1
		}
	}

	// target was not present in the array
	return -1
}

func main() {
	arr := []int{2, 3, 4, 10, 40}
	target := 10

	result := binarySearch(arr, target)
	if result != -1 {
		fmt.Printf("Element is present at index %d\n", result)
	} else {
		fmt.Println("Element is not present in array")
	}
}

