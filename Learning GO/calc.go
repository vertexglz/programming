package main

import (
	"errors"
	"fmt"
)

// Все функции объявляются ВНЕ main()
func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль запрещено")
	}
	return a / b, nil
}

func main() {
	var history []string
	for {
		var num1, num2 float64
		var operator string

		fmt.Print("Введите первое число: ")
		fmt.Scanln(&num1)

		fmt.Print("Введите операцию: ")
		fmt.Scanln(&operator)

		fmt.Print("Введите второе число: ")
		fmt.Scanln(&num2)

		var res float64
		var err error
		calcError := false

		switch operator {
		case "+":
			res = add(num1, num2)
		case "-":
			res = subtract(num1, num2)
		case "*":
			res = multiply(num1, num2)
		case "/":
			res, err = divide(num1, num2)
			if err != nil {
				fmt.Println("Ошибка:", err)
				calcError = true
			}

		default:
			fmt.Println("Ошибка: неизвестная операция!")
			calcError = true
		}

		if !calcError {
			fmt.Printf("Результат %.2f\n", res)
			record := fmt.Sprintf("%2.f %s %2.f = %2.f", num1, operator, num2, res)
			history = append(history, record)
		}

		fmt.Print("Если хотите выйти то напишите exit, история вычислений - history если продолжить то нажмите Enter: ")
		var choise string
		fmt.Scanln(&choise)
		if choise == "exit" {
			break
		}
		if choise == "history" {
			fmt.Println("---История операций---")
			if len(history) == 0 {
				fmt.Println("История пока что пуста")
			}
			for i, item := range history {
				fmt.Printf("%d: %s\n", i+1, item)
			}
		}
	}
}
