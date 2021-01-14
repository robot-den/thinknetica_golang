module pkg/api

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	pkg/auth v1.0.0
	pkg/model v1.0.0
)

replace pkg/auth => ./../auth
replace pkg/model => ../model

go 1.15
