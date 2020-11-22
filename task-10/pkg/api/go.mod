module pkg/api

go 1.14

replace (
	pkg/crawler => ../crawler
	pkg/engine => ../engine
	pkg/index => ../index
	pkg/index/hash => ../index/hash
	pkg/storage => ../storage
	pkg/storage/memstore => ../storage/memstore
)

require (
	github.com/gorilla/mux v1.8.0
	github.com/prometheus/client_golang v1.8.0
	pkg/engine v0.0.0-00010101000000-000000000000
	pkg/index/hash v0.0.0-00010101000000-000000000000
	pkg/storage/memstore v0.0.0-00010101000000-000000000000
)
