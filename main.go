package main

import "fmt"

func main() {
	arr := []int{2, 5, 10, 14, 23, 42, 51, 57, 65}
	target := 10
	low := 0

	high := len(arr) - 1
	index := 0
	//	mid := midElement(low, high)
	for low <= high {
		fmt.Println("low:", low)
		fmt.Println("high:", high)
		mid := midElement(low, high)
		fmt.Println("mid:", mid)
		if arr[mid] == target {
			fmt.Println("mid element", mid, arr[mid], target)
			index = mid
			break
		} else if arr[mid] < target {
			fmt.Println("mid element", mid, arr[mid], target)
			low = mid + 1

		} else {
			fmt.Println("mid element", mid, arr[mid], target)
			high = mid - 1

		}
	}

	fmt.Println("index element is: ", index)
}

func midElement(low, high int) int {

	mid := (low + high) / 2

	return mid

}
