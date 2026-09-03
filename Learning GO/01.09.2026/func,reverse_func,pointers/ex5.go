package main

import "fmt"

func factorial(n int, result *int) {
	if n <= 1 {
		return
	}
	*result *= n
	factorial(n-1, result)

}

func main() {
	res := 1
	factorial(5, &res)
	fmt.Println(res)

}
