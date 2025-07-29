package main

import "fmt"

// Структура для хранения данных о сотруднике
type Employee struct {
	lastName  string
	firstName string
	age       int
	position  string
	salary    int
}

// Интерфейс для вывода информации
type Displayable interface {
	Display()
}

// Реализация метода Display для Employee
func (e Employee) Display() {
	fmt.Printf("Фамилия: %s, Имя: %s, Возраст: %d, Должность: %s, Зарплата: %d рублей\n", e.lastName, e.firstName, e.age, e.position, e.salary)
}

// Функция фильтрации сотрудников по возрасту и зарплате
func FilterEmployees(employees []Employee, minAge int, minSalary int) []Employee {
	filtered := make([]Employee, 0)
	for _, emp := range employees {
		if emp.age >= minAge && emp.salary >= minSalary {
			filtered = append(filtered, emp)
		}
	}
	return filtered
}

func main() {
	// Инициализация списка сотрудников
	employees := []Employee{
		{"Иванов", "Иван", 30, "Инженер", 50000},
		{"Петров", "Петр", 40, "Менеджер", 60000},
		{"Васильев", "Василий", 25, "Стажер", 35000},
		{"Смирнова", "Анна", 35, "Директор", 80000},
	}

	// Параметры фильтрации
	minAge := 30
	minSalary := 50000

	// Фильтрация и вывод
	filteredEmployees := FilterEmployees(employees, minAge, minSalary)
	fmt.Println("Список сотрудников после фильтрации:")
	for _, emp := range filteredEmployees {
		emp.Display()
	}
}