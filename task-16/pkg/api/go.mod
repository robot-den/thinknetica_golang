module pkg/api

require (
	github.com/gorilla/mux v1.8.0
	pkg/engine v1.0.0
	pkg/index v1.0.0
	pkg/index/hash v1.0.0
	pkg/model v1.0.0
	pkg/storage v1.0.0
	pkg/storage/memory v1.0.0
)

replace pkg/engine => ../engine
replace pkg/model => ../../pkg/model
replace pkg/index => ../../pkg/index
replace pkg/index/hash => ../../pkg/index/hash
replace pkg/storage => ../../pkg/storage
replace pkg/storage/memory => ../../pkg/storage/memory

go 1.15
