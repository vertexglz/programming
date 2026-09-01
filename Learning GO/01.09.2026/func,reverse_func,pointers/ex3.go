package main

import "fmt"

// newWallet возвращает АНОНИМНУЮ ФУНКЦИЮ с сигнатурой func(amount int) (int, bool)
func newWallet(initialBalance int) func(amount int) (int, bool) {
	// Возвращаем функцию прямо через return
	return func(amount int) (int, bool) {
		if initialBalance >= amount {
			initialBalance -= amount // Меняем состояние, захваченное замыканием
			return initialBalance, true
		}
		return initialBalance, false // Если денег не хватает, баланс не меняем
	}
}

func main() {
	// 1. Создаем кошелек с балансом 1000.
	// Переменная myWallet теперь сама является ФУНКЦИЕЙ!
	myWallet := newWallet(1000)

	// 2. Вызываем myWallet как обычную функцию, передавая только amount
	balance, ok := myWallet(1200)
	fmt.Printf("Сняли 1200: Успех=%t, Остаток=%d\n", ok, balance) // true, 700

	balance, ok = myWallet(800)
	fmt.Printf("Сняли 800: Успех=%t, Остаток=%d\n", ok, balance) // false, 700

	balance, ok = myWallet(200)
	fmt.Printf("Сняли 200: Успех=%t, Остаток=%d\n", ok, balance) // true, 500
}
