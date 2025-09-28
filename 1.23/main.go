package main

import "fmt"

func main() {
	arr := []string{"a", "b", "c", "d", "e"}
    
	fmt.Println("Before:", arr)
    arr = removeEl(arr, 2)
    fmt.Println("After: ", arr)
}

func removeEl(arr []string, i int) []string {
	copy(arr[i:], arr[i+1:])
	arr[len(arr)-1] = ""
	return arr[:len(arr)-1]
}