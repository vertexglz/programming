Проходим псевдонимы и структуры
Вот теория:
1. Пользовательские типы (Defined Types)
Позволяют создавать собственные типы на базе существующих для добавления контекста и строгой типизации.

Go
type Celsius float64
type UserID int

var temp Celsius = 36.6
// var f float64 = temp // Ошибка! Нужна явное приведение типов: float64(temp)
2. Структуры (Structs)
Структура — это составной тип данных, объединяющий логически связанные поля разных типов.

Go
// Объявление типа структуры
type Person struct {
    Name string
    Age  int
}

// Способы инициализации
p1 := Person{Name: "Алексей", Age: 25} // Инициализация по именам полей (рекомендуется)
p2 := Person{"Мария", 30}              // Инициализация по порядку полей
var p3 Person                          // Нулевые значения (Name: "", Age: 0)
3. Доступ к полям и указатели на структуры
Доступ к полям осуществляется через точку (.). При работе с указателями Go автоматически разыменовывает их для доступа к полям — синтаксис остается простым.

Go
p := Person{Name: "Иван", Age: 20}
p.Age = 21 // Изменение поля напрямую

// Указатель на структуру
pPtr := &p
pPtr.Age = 22 // В Go это эквивалентно (*pPtr).Age = 22 (авто-разыменование)

// Создание структуры сразу в виде указателя
p2Ptr := &Person{Name: "Ольга", Age: 28}
4. Анонимные структуры (Anonymous Structs)
Используются для одноразовых задач (например, парсинг конфигурации, JSON или быстрая группировка данных в тестах), когда не нужно объявлять глобальный тип через type.

Go
user := struct {
    Username string
    IsAdmin  bool
}{
    Username: "admin",
    IsAdmin:  true,
}
5. Вложенность и встраивание структур
Обычное вложение (Explicit nesting):
Go
type Address struct {
    City string
}

type User struct {
    Name    string
    Addr Address // Явное поле с типом Address
}

// Обращение: u.Addr.City
Анонимное встраивание (Embedded / Promoted Fields):
Поля встроенной структуры «продвигаются» на верхний уровень.

Go
type Contact struct {
    Email string
}

type Employee struct {
    Name string
    Contact // Анонимное встраивание (без имени поля)
}

emp := Employee{Name: "Петр", Contact: Contact{Email: "test@example.com"}}
emp.Email = "new@example.com" // Прямой доступ к полю встроенной структуры (promoted field)
Отличный задел для конспекта! Текст полностью готов для сохранения в .txt или README.md.

Когда будешь готов — скидывай решения первых 3 лёгких задач (про Celsius, Student и Wallet), разберём их вместе.
Задачи на сегодня:
Задача 1. Кастомные типы (Defined Types)Создай собственный тип Celsius на базе float64.Напиши функцию ToFahrenheit(c Celsius) float64, которая переводит температуру из Цельсий в Фаренгейты по формуле: $F = C \times 1.8 + 32$.Цель: Закрепить объявление пользовательских типов и передачу их в функции.
Задача 2. Базовая структура и проверкаСоздай структуру Student со следующими полями:Name (string)Age (int)AverageGrade (float64)Напиши функцию IsHonorsStudent(s Student) bool, которая возвращает true, если AverageGrade >= 4.5, и false в противном случае.
Задача 3. Изменение структуры через указательСоздай структуру Wallet с полем Balance (int).Напиши функцию Deposit(w *Wallet, amount int), которая увеличивает Balance на amount.Ограничение: Функция ничего не возвращает (void), изменение баланса должно происходить строго через указатель на структуру.Задачи среднего уровня (Вложенность, встраивание и коллекции)
Задача 4. Вложенные структуры (Nested Structs)Создай две структуры:Address с полями City (string), Street (string), House (int).Employee с полями Name (string), Position (string), WorkAddress (Address).Напиши функцию GetEmployeeCard(e Employee) string, которая возвращает красиво форматированную строку вида:"Иван (Developer) - г. Москва, ул. Ленина, д. 10"
Задача 5. Встраивание структур (Embedded Structs / Promoted Fields)Создай структуру ContactInfo с полями Email (string) и Phone (string).Создай структуру User, которая анонимно встраивает ContactInfo (без имени поля, только тип) и имеет собственные поля ID (int) и Username (string).Напиши функцию UpdateEmail(u *User, newEmail string), которая меняет email пользователя.Цель: Проверить, как работает прямой доступ к полям встроенной структуры через обращения вида u.Email (promoted fields).
Задача 6. Анализ массива структурСоздай структуру Product с полями Name (string), Price (float64), InStock (bool).Напиши функцию CalculateInventory(products []Product) (totalValue float64, availableCount int):Принимает слайс товаров []Product.Возвращает общую стоимость всех товаров, которые есть в наличии (InStock == true), и их общее количество.