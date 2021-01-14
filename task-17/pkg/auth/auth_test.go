package auth

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"pkg/model"
	"reflect"
	"testing"
)

func TestAuth_Authorize(t *testing.T) {
	auth := New()
	creds := model.Credentials{
		Login:    "admin",
		Password: "strong",
	}
	token_string, err := auth.Authorize(creds)
	if err != nil {
		t.Errorf("auth.Authorize(); err = %v; want %v", err, nil)
	}

	// https://godoc.org/github.com/dgrijalva/jwt-go#ParseWithClaims
	token, err := jwt.ParseWithClaims(token_string, &model.RightsClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})
	if err != nil {
		t.Errorf("jwt.Parse(); err = %v; want %v", err, nil)
	}

	if !token.Valid {
		t.Errorf("token.Valid = %v, want %v", false, true)
	}

	claims, ok := token.Claims.(*model.RightsClaims)
	if !ok {
		t.Errorf("token.Claims.(jwt.RightsClaims); ok = %v, want %v", false, true)
	}

	got := claims.Rights
	want := []string{"create", "read", "update", "delete"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("reflect.DeepEqual(got, want) = %v; want %v", false, true)
	}
}
