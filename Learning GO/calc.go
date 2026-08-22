package main

import "fmt"

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&num1)

	fmt.Print("Введите операцию: ")
	fmt.Scanln(&operator)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&num2)

	switch operator {
	case "+":
		fmt.Printf("Результат: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Результат: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Результат: %.2f\n", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Printf("Ошибка деления на ноль")
		} else {
			fmt.Printf("Результат: %.2f\n", num1/num2)
		}
	default:
		fmt.Println("Ошибка: неизвестная операция!")
	}
}
