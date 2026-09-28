package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	other := make([]int, len(numbers))
	copy(other, numbers)
	other = append(other[:2], other[3:]...)
	other = append(other, 0)
	copy(other[3:], other[2:])
	other[2] = 99
	other[0] = 999
	fmt.Println(numbers)
	fmt.Println(other)

}
