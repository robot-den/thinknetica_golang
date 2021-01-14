// package auth реализует сервис аутентификации и авторизации
package auth

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"pkg/model"
	"time"
)

// secret это ключ которым подписывается JWT-токен
var secret = []byte("very-secret-sign")

// passwords это "база пользователей", содержащая пары логин/пароль
var passwords = map[string]string{
	"admin": "$2a$10$LkPvXflA7CNz2r0G1CKxj.lU1IJ07NEsbTHSm10GkFU/0MZt4P2d6",
	"guest": "$2a$10$.kIUjKRBiHHcoYc.my0EzO/v9iVvWbTHcNXKyiXrlEfj07z6pa1KG",
}

// accessRights это "база прав", содержит информацию о том какие действия над ресурсами может выполнять пользователь
var accessRights = map[string][]string{
	"admin": []string{"create", "read", "update", "delete"},
	"guest": []string{"read"},
}

// Auth представляет собой объект реализующий методы-обработчики API
type Auth struct{}

// New создает новый объект Auth
func New() *Auth {
	return &Auth{}
}

// Authorize аутентифицирует и авторизует пользователя по паре логин/пароль
func (a *Auth) Authorize(credentials model.Credentials) (string, error) {
	err := a.authenticate(credentials)
	if err != nil {
		return "", err
	}

	token := a.buildToken(credentials)

	return token, nil
}

func (a *Auth) authenticate(credentials model.Credentials) error {
	error := fmt.Errorf("пользователь с таким логином или паролем не существует")

	hashedPassword, ok := passwords[credentials.Login]
	if !ok {
		return error
	}

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(credentials.Password))
	if err != nil {
		return error
	}

	return nil
}

func (a *Auth) buildToken(credentials model.Credentials) string {
	rights := accessRights[credentials.Login]

	claims := model.RightsClaims{
		Rights: rights,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(secret)
	return tokenString
}
