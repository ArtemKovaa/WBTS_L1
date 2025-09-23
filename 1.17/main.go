package main

import "log"

func main() {
	arr := []int{-17, -15, -9, 0, 1, 6, 7, 10, 117, 118, 1110, 2222}
	el := 2222

	log.Printf("Element %d was found by index %d with binarySearchFor", el, binarySearchFor(arr, el))
	log.Printf("Element %d was found by index %d with binarySearchRecursion", el, binarySearchRecursion(arr, el, 0, len(arr)-1))
}

func binarySearchFor(slice []int, el int) int {
	left, right := 0, len(slice)-1

	for left <= right {
		mid := (left + right) / 2

		if slice[mid] == el {
			return mid
		} else if slice[mid] < el {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func binarySearchRecursion(slice []int, el, left, right int) int {
	if left > right {
		return -1
	}

	mid := (left + right) / 2

	if slice[mid] == el {
		return mid
	} else if slice[mid] < el {
		return binarySearchRecursion(slice, el, mid+1, right)
	}
	return binarySearchRecursion(slice, el, left, mid-1)
}
