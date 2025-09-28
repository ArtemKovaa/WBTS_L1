package main

import "log"

func main() {
	arr := []int{8, 0, 11, 17, -100, -1, -4, -3, 82, 11, 100, 22, 1001, 1, -2}
	
	quickSort(arr)

	log.Println(arr)
}


func quickSort(arr []int) {
	sort(arr, 0, len(arr)-1)
}

func sort(arr []int, low, high int) {
    if low < high {
		pivotIndex := partition(arr, low, high)
		sort(arr, low, pivotIndex-1)
		sort(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
    pivot := arr[high]
    i := low - 1
    for j := low; j < high; j++ {
        if arr[j] < pivot {
            i++
            arr[i], arr[j] = arr[j], arr[i]
        }
    }
    arr[i+1], arr[high] = arr[high], arr[i+1]
    return i + 1
}
