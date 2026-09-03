package main

import "fmt"

func sumDigit(n int) int {
	if n < 0 {
		n = -n
	}

	if n < 10 {
		return n
	}
	return (n % 10) + sumDigit(n/10)
}

func main() {
	fmt.Println(sumDigit(12340))
	fmt.Println(sumDigit(-53235534))

}
