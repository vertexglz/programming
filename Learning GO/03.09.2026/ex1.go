package main

import "fmt"

type Celsius float64

func ToFarenheit(c Celsius) float64 {
	return float64(c)*1.8 + 32
}

func main() {
	var n Celsius
	n = 36.6

	fmt.Printf("%.2f\n", ToFarenheit(n))

}
