// Package problem1 реализует решение для задания №1
package problem1

// Employee описывает сотрудника организации
type Employee struct {
	age int
}

// Age возвращает возраст сотрудника
func (e Employee) Age() int {
	return e.age
}

// Customer описывает клиента организации
type Customer struct {
	age int
}

// Age возвращает возраст клиента
func (c Customer) Age() int {
	return c.age
}

type user interface {
	Age() int
}

// MaxAge возвращает возраст самого старшего пользователя
func MaxAge(users ...user) int {
	var maxAge int

	for _, u := range users {
		if maxAge < u.Age() {
			maxAge = u.Age()
		}
	}
	return maxAge
}
