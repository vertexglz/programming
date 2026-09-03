package main

import "fmt"

func outer() func() {
	var n int = 5
	inner := func() {
		n += 1
		fmt.Println(n)
	}
	return inner
}

func main() {
	fn := outer()
	fn()
	fn()
	fn()
}
