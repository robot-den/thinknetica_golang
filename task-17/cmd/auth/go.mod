module cmd/auth

require (
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	pkg/api v1.0.0
	pkg/auth v1.0.0
	pkg/model v1.0.0
)

replace pkg/api => ../../pkg/api
replace pkg/auth => ./../../pkg/auth
replace pkg/model => ./../../pkg/model

go 1.15
