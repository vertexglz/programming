package main

import "fmt"

type Adress struct {
	City, Street string
	House        int
}

type Employee struct {
	Name, Position string
	Work_adress    Adress
}

func GetEmployeeCard(e Employee) string {
	// 1. Добавлен return
	// 2. Использование экземпляра 'e' вместо типа 'Employee'
	// 3. %d для целого числа House
	return fmt.Sprintf("%s (%s) - г. %s, ул. %s, д. %d",
		e.Name,
		e.Position,
		e.Work_adress.City,
		e.Work_adress.Street,
		e.Work_adress.House,
	)
}

func main() {
	// 4. Указание типа при инициализации
	MTS := Adress{"Глазов", "Пл.Свободы", 7}
	pavel := Employee{"Павел", "Менеджер по продажам", MTS}

	// Печатаем результат возврата функции
	fmt.Println(GetEmployeeCard(pavel))
}
