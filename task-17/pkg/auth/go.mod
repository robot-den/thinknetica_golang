module pkg/auth

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	pkg/model v1.0.0
)

replace pkg/model => ../model

go 1.15
