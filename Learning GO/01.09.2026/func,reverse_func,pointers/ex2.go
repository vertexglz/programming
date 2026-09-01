package main

import "fmt"

func filterSlice(nums []int, check func(int) bool) []int {
	result := []int{}
	for _, value := range nums {
		if check(value) {
			result = append(result, value)
		}

	}
	return result
}

func main() {
	slice := []int{-1, 2, -3, 4, -5, 6, -7}
	f := func(n int) bool {
		if n > 0 {
			return true
		}
		return false
	}
	fmt.Println(filterSlice(slice, f))

	fmt.Println(filterSlice(slice, func(n int) bool { // можно объявлять анонимную функцию прямо в выводе
		return n%3 == 0 && n > 0
	}))

}
