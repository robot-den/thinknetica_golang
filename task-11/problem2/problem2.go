// Package problem2 реализует решение для задания №2
package problem2

import "fmt"

// Employee описывает сотрудника организации
type Employee struct {
	Age int
}

// Customer описывает клиента организации
type Customer struct {
	Age int
}

// MaxAge возвращает самого старшего пользователя
func MaxAge(users ...interface{}) (interface{}, error) {
	var oldest interface{}
	var maxAge int

	for _, u := range users {
		var age int

		switch v := u.(type) {
		case Employee:
			age = v.Age
		case Customer:
			age = v.Age
		default:
			return nil, fmt.Errorf("can't determine age attribute of %T", v)
		}

		if age > maxAge {
			maxAge = age
			oldest = u
		}
	}
	return oldest, nil
}
