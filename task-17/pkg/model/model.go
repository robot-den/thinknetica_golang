// package model содержит типы данных, совместно используемые в разных пакетах приложения
package model

import (
	jwt "github.com/dgrijalva/jwt-go"
)

// Credentials представляет собой учетные данные пользователя
type Credentials struct {
	Login    string
	Password string
}

// RightClaims представляет собой объект содержащий набор прав для проведения авторизации
type RightsClaims struct {
	Rights []string `json:"rts"`
	jwt.StandardClaims
}

// Valid позволяет провалидировать содержимое объекта. Метод необходим чтобы можно было использовать RightsClaims в jwt
func (r RightsClaims) Valid() error {
	return nil
}
