package main

// Determine if an array is in ascending order or not
import "fmt"

func isAscending(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 5, 3, 4}
	arr3 := []int{5, 4, 3, 2, 1}

	fmt.Println(isAscending(arr1)) // prints "true"
	fmt.Println(isAscending(arr2)) // prints "false"
	fmt.Println(isAscending(arr3)) // prints "false"
}